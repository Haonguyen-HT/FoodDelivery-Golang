package ginrestaurant

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	restaurantbusiness "FoodDelivery/module/restaurant/business"
	retaurantstorage "FoodDelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindRestaurant(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := retaurantstorage.NewSQLStore(db)
		biz := restaurantbusiness.NewFindRestaurantBusiness(store)

		result, err := biz.FindRestaurant(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(&result))

	}
}