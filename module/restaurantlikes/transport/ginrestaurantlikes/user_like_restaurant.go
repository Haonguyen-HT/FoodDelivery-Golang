package ginrestaurantlikes

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	restaurantstorage "FoodDelivery/module/restaurant/storage"
	rstlikebiz "FoodDelivery/module/restaurantlikes/biz"
	restaurantlikesmodel "FoodDelivery/module/restaurantlikes/model"
	restaurantlikestorage "FoodDelivery/module/restaurantlikes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserLikeRestaurant(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(err)
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data = restaurantlikesmodel.UserLike{
			RestaurantId: id,
			UserId:       requester.GetUserID(),
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantlikestorage.NewSQLStore(db)
		icrLikeStore := restaurantstorage.NewSQLStore(db)

		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, icrLikeStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse("like restaurant successfully"))

	}
}