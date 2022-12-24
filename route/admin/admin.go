package admin

import (
	"FoodDelivery/components/appcontext"
	"FoodDelivery/middleware"
	"FoodDelivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func AdminRoute(appContext appcontext.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin",
		middleware.RequireAuth(appContext),
		middleware.RoleChecker(appContext,
			"admin", "mod"))

	admin.GET("/profile", ginuser.Profile(appContext))
}