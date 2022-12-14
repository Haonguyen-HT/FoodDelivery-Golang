package subscriber

import (
	"FoodDelivery/components/appcontext"
	restaurantstorage "FoodDelivery/module/restaurant/storage"
	"FoodDelivery/pubsub"
	"context"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(
	appCtx appcontext.AppContext) consumerJob {

	return consumerJob{
		Title: "Increase like count after user like restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}
func RealtimeAfterUserLikeRestaurant(
	appCtx appcontext.AppContext) consumerJob {

	return consumerJob{
		Title: "Realtime emit after user like restaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			appCtx.GetRealtimeEngine().EmitToUser(likeData.GetUserId(), string(message.Chanel()), likeData)

			return nil
		},
	}

}