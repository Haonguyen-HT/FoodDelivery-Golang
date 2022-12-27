package skio

import (
	"FoodDelivery/components/tokenprovider/jwt"
	userstore "FoodDelivery/module/user/store"
	"FoodDelivery/module/user/transport/skuser"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"gorm.io/gorm"
	"log"
	"sync"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	GetRealtimeEngine() *RtEngine
}

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(ctx AppContext, engine *gin.Engine) error
}

type RtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *RtEngine {
	return &RtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (engine *RtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()

	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}

	engine.locker.Unlock()
}

func (engine *RtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	return engine.storage[userId]
}

func (engine *RtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()

	defer engine.locker.Unlock()

	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
}

func (engine *RtEngine) UserSockets(userId int) []AppSocket {
	var sockets []AppSocket

	if scks, ok := engine.storage[userId]; ok {
		return scks
	}

	return sockets
}

func (engine *RtEngine) EmitToRoom(room, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (engine *RtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)

	for _, s := range sockets {
		s.Emit(key, data)
	}
	return nil
}

func (engine *RtEngine) Run(appCtx AppContext, r *gin.Engine) {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	engine.server = server

	server.OnConnect("/", func(s socketio.Conn) error {
		//s.SetContext("")
		s.Join("test")
		fmt.Print("connected:", s.ID(), " IP:", s.RemoteAddr())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, err error) {
		fmt.Print("meet error", err)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Print("close", reason)
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {

		db := appCtx.GetMainDBConnection()
		store := userstore.NewSQLStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			s.Emit("authenticate", err.Error())
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			s.Emit("authenticate", err.Error())
			return
		}

		if user.Status == 0 {
			s.Emit("authenticate", errors.New("you has been banned/deleted"))
			return
		}

		appSck := NewAppSocket(s, user)
		engine.saveAppSocket(user.Id, appSck)

		server.OnEvent("/", "UserUpdateLocation", skuser.OnUserUpdateLocation(appCtx, user))

		s.Emit("authenticate", user)
	})

	go server.Serve()

	go func() {
		defer server.Close()
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
}