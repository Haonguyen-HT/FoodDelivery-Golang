package client

import (
	"FoodDelivery/components/appcontext"
	"FoodDelivery/middleware"
	"FoodDelivery/module/restaurant/transport/ginrestaurant"
	"FoodDelivery/module/restaurantlikes/transport/ginrestaurantlikes"
	"github.com/gin-gonic/gin"
)

func RestaurantRoute(appContext appcontext.AppContext, v1 *gin.RouterGroup) {
	restaurants := v1.Group("/restaurants")

	// POST /restaurants

	restaurants.POST("", middleware.RequireAuth(appContext), ginrestaurant.CreateRestaurant(appContext))
	restaurants.POST("/:id/like", middleware.RequireAuth(appContext), ginrestaurantlikes.UserLikeRestaurant(appContext))
	restaurants.DELETE("/:id/dislike", middleware.RequireAuth(appContext), ginrestaurantlikes.UserDislikeRestaurant(appContext))

	// GET all restaurants

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	// GET restaurants by ID
	restaurants.GET("/:id", ginrestaurant.FindRestaurant(appContext))

	// DELETE Restaurant by id

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	// UPDATE Restaurant by id

	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))
}