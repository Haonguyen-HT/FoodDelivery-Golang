package main

import (
	"FoodDelivery/components/appcontext"
	"FoodDelivery/middleware"
	"FoodDelivery/pubsub/localPb"
	"FoodDelivery/route/admin"
	"FoodDelivery/route/client"
	"FoodDelivery/route/user"
	"FoodDelivery/subscriber"
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
	secretKey := os.Getenv("SECRET_KEY")

	//s3BucketName := os.Getenv("S3BucketName")
	//s3Region := os.Getenv("S3Region")
	//s3APIKEY := os.Getenv("S3APIKEY")
	//s3SecretKey := os.Getenv("S3SecretKey")
	//s3Domain := os.Getenv("S3Domain")
	//
	//s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKEY, s3SecretKey, s3Domain)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	if err != nil {
		log.Fatalln(err)
	}

	ps := localPb.NewPubSub()

	appContext := appcontext.NewAppContext(db, nil, secretKey, ps)
	// Set up subscriber

	subscriber.NewEngine(appContext).Start()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true

	r := gin.Default()
	r.Use(cors.New(config))
	r.Use(middleware.Recover(appContext))
	r.Static("/static", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	admin.AdminRoute(appContext, v1)
	user.UserRoute(appContext, v1)
	client.RestaurantRoute(appContext, v1)

	r.Run(":3001").Error() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}