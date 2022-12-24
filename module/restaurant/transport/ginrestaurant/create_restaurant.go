package ginrestaurant

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	restaurantbusiness "FoodDelivery/module/restaurant/business"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	retaurantstorage "FoodDelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserID = requester.GetUserID()

		store := retaurantstorage.NewSQLStore(db)
		biz := restaurantbusiness.NewCreateRestaurantBusiness(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.UID.String()))

	}
}