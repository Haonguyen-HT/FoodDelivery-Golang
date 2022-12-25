package subscriber

import (
	"FoodDelivery/components/appcontext"
	restaurantstorage "FoodDelivery/module/restaurant/storage"
	"FoodDelivery/pubsub"
	"context"
)

type HasRestaurantId interface {
	GetRestaurantId() int
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