package ginrestaurantlikes

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	rstlikebiz "FoodDelivery/module/restaurantlikes/biz"
	restaurantlikestorage "FoodDelivery/module/restaurantlikes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserDislikeRestaurant(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(err)
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		UserId := requester.GetUserID()

		store := restaurantlikestorage.NewSQLStore(db)
		//dcrStore := restaurantstorage.NewSQLStore(db)

		biz := rstlikebiz.NewUserDislikeRestaurantBiz(store, ctx.GetPubSub())

		if err := biz.DislikeRestaurant(c.Request.Context(), UserId, id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(true))

	}
}