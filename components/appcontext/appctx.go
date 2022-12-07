package appcontext

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (ctx appContext) GetMainDBConnection() *gorm.DB {
	return ctx.db
}