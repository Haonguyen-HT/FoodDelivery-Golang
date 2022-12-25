package subscriber

import (
	"FoodDelivery/components/appcontext"
	restaurantstorage "FoodDelivery/module/restaurant/storage"
	"FoodDelivery/pubsub"
	"context"
	"log"
)

func DecreaseLikeCountAfterUserLikeRestaurant(
	appCtx appcontext.AppContext) consumerJob {

	return consumerJob{
		Title: "Decrease like count after user dislike restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}

func PushNotificationWhenUserDisLikeRestaurant(
	appCtx appcontext.AppContext) consumerJob {

	return consumerJob{
		Title: "Push notification when user dislike restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push notification when user dislike restaurant id:", likeData.GetRestaurantId())

			return nil
		},
	}

}