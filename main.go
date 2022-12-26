package main

import (
	"FoodDelivery/components/appcontext"
	"FoodDelivery/components/tokenprovider/jwt"
	"FoodDelivery/middleware"
	userstore "FoodDelivery/module/user/store"
	"FoodDelivery/pubsub/localPb"
	"FoodDelivery/route/admin"
	"FoodDelivery/route/client"
	"FoodDelivery/route/user"
	"FoodDelivery/subscriber"
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	secretKey := os.Getenv("SECRET_KEY")

	//s3BucketName := os.Getenv("S3BucketName")
	//s3Region := os.Getenv("S3Region")
	//s3APIKEY := os.Getenv("S3APIKEY")
	//s3SecretKey := os.Getenv("S3SecretKey")
	//s3Domain := os.Getenv("S3Domain")
	//
	//s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKEY, s3SecretKey, s3Domain)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	if err != nil {
		log.Fatalln(err)
	}

	ps := localPb.NewPubSub()

	appContext := appcontext.NewAppContext(db, nil, secretKey, ps)
	// Set up subscriber

	subscriber.NewEngine(appContext).Start()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true

	r := gin.Default()
	r.Use(cors.New(config))
	r.Use(middleware.Recover(appContext))
	r.Static("/static", "./static")
	r.StaticFile("/demo", "./demosocket.html")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	admin.AdminRoute(appContext, v1)
	user.UserRoute(appContext, v1)
	client.RestaurantRoute(appContext, v1)

	startSocketIOServer(r, appContext)
	r.Run(":3001").Error() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func startSocketIOServer(engine *gin.Engine, appCtx appcontext.AppContext) {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Print("connected:", s.ID(), " IP:", s.RemoteAddr())

		s.Join("Shipper")

		return nil
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

		s.Emit("authenticate", user)
	})

	server.OnEvent("/", "test", func(s socketio.Conn, msg interface{}) {
		log.Println("message test", msg)
	})

	server.OnError("/", func(s socketio.Conn, err error) {
		fmt.Print("meet error", err)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Print("close", reason)
	})

	go server.Serve()

	go func() {
		defer server.Close()
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	engine.GET("/socket.io/*any", gin.WrapH(server))
	engine.POST("/socket.io/*any", gin.WrapH(server))
}