package main

import (
	"FoodDelivery/components/appcontext"
	"FoodDelivery/middleware"
	"FoodDelivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	if err != nil {
		log.Fatalln(err)
	}

	appContext := appcontext.NewAppContext(db)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true

	r := gin.Default()
	r.Use(cors.New(config))
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	// POST /restaurants

	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	// GET all restaurants

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	// GET restaurants by ID
	restaurants.GET("/:id", ginrestaurant.FindRestaurant(appContext))

	// DELETE Restaurant by id

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	// UPDATE Restaurant by id

	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))

	r.Run(":3001").Error() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}