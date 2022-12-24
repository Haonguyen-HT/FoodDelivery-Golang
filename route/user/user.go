package user

import (
	"FoodDelivery/components/appcontext"
	"FoodDelivery/middleware"
	"FoodDelivery/module/upload/transport/ginupload"
	"FoodDelivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func UserRoute(appContext appcontext.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/login", ginuser.Login(appContext))
	v1.POST("/profile", middleware.RequireAuth(appContext), ginuser.Profile(appContext))

	v1.POST("/upload", ginupload.UploadImage(appContext))
}