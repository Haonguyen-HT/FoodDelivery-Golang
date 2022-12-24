package rstlikebiz

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	restaurantlikesmodel "FoodDelivery/module/restaurantlikes/model"
	"context"
	"fmt"
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
	IncreaseLikeCount(context context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store        UserLikeRestaurantStore
	icrLikeCount IncreaseLikeCount
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, count IncreaseLikeCount) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, icrLikeCount: count}
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

	go func() {
		defer common.AppRecover()
		if err := biz.icrLikeCount.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}