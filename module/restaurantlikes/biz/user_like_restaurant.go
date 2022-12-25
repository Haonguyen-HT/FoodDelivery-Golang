package rstlikebiz

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	restaurantlikesmodel "FoodDelivery/module/restaurantlikes/model"
	"FoodDelivery/pubsub"
	"context"
	"log"
)

type UserLikeRestaurantStore interface {
	CreateLikeRestaurant(context context.Context, data *restaurantlikesmodel.UserLike) error
}

type IncreaseLikeCount interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type userLikeRestaurantBiz struct {
	store        UserLikeRestaurantStore
	icrLikeCount IncreaseLikeCount
	ps           pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	count IncreaseLikeCount,
	ps pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, icrLikeCount: count, ps: ps}
}

func (biz userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikesmodel.UserLike) error {
	_, err := biz.icrLikeCount.FindDataWithCondition(
		ctx,
		map[string]interface{}{
			"id": data.RestaurantId,
		})

	if err != nil {
		return restaurantlikesmodel.ErrCannotLikeRestaurant(err)
	}

	if err := biz.store.CreateLikeRestaurant(ctx, data); err != nil {
		return restaurantlikesmodel.ErrCannotLikeRestaurant(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	return nil
}