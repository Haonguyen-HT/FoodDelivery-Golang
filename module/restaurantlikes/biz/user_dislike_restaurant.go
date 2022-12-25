package rstlikebiz

import (
	"FoodDelivery/common"
	restaurantlikesmodel "FoodDelivery/module/restaurantlikes/model"
	"FoodDelivery/pubsub"
	"context"
	"log"
)

type UserDislikeRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantlikesmodel.Like, error)
	Delete(context context.Context, userId, restaurantId int) error
}

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
	ps    pubsub.Pubsub
}

func NewUserDislikeRestaurantBiz(
	store UserDislikeRestaurantStore,
	ps pubsub.Pubsub) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, ps: ps}
}

func (biz userDislikeRestaurantBiz) DislikeRestaurant(
	ctx context.Context,
	userId, restaurantId int,
) error {
	_, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"restaurant_id": restaurantId, "user_id": userId})
	if err != nil {
		return restaurantlikesmodel.ErrCannotDisLikeRestaurant(err)
	}

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikesmodel.ErrCannotDisLikeRestaurant(err)
	}

	if err := biz.ps.Publish(
		ctx,
		common.TopicUserDisLikeRestaurant,
		pubsub.NewMessage(&restaurantlikesmodel.UserLike{RestaurantId: restaurantId})); err != nil {
		log.Println(err)
	}

	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.DcrStore.DecreaseLikeCount(ctx, restaurantId)
	//})
	//
	//if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	//	log.Println(err)
	//}

	return nil
}