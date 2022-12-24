package rstlikebiz

import (
	"FoodDelivery/common"
	restaurantlikesmodel "FoodDelivery/module/restaurantlikes/model"
	"context"
	"fmt"
)

type UserDislikeRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantlikesmodel.Like, error)
	Delete(context context.Context, userId, restaurantId int) error
}

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(context context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store    UserDislikeRestaurantStore
	DcrStore DecreaseLikeCountStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, countStore DecreaseLikeCountStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, DcrStore: countStore}
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

	go func() {
		defer common.AppRecover()
		if err := biz.DcrStore.DecreaseLikeCount(ctx, restaurantId); err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}