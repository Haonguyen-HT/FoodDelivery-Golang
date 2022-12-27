package appcontext

import (
	"FoodDelivery/components/uploadprovider"
	"FoodDelivery/pubsub"
	"FoodDelivery/skio"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
	GetRealtimeEngine() *skio.RtEngine
}

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
	rtEngine       *skio.RtEngine
}

func NewAppContext(
	db *gorm.DB,
	uploadprovider *uploadprovider.UploadProvider,
	secretKey string,
	ps pubsub.Pubsub) *appContext {
	return &appContext{db: db, uploadProvider: nil, secretKey: secretKey, ps: ps}
}

func (ctx *appContext) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appContext) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appContext) GetPubSub() pubsub.Pubsub {
	return ctx.ps
}
func (ctx *appContext) GetRealtimeEngine() *skio.RtEngine {
	return ctx.rtEngine
}

func (ctx *appContext) SetRealtimeEngine(rt *skio.RtEngine) {
	ctx.rtEngine = rt
}