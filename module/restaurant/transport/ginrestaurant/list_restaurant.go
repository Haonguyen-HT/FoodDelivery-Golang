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

func ListRestaurant(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		pagingData.Fulfill()

		var result []restaurantmodel.Restaurant

		store := retaurantstorage.NewSQLStore(db)
		biz := restaurantbusiness.NewListRestaurantBusiness(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}