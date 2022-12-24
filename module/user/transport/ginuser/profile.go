package ginuser

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u := context.MustGet(common.CurrentUser).(common.Requester)
		fmt.Println(u)
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}