package restaurantbusiness

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

type FindRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBusiness struct {
	store FindRestaurantStore
}

func NewFindRestaurantBusiness(store FindRestaurantStore) *findRestaurantBusiness {
	return &findRestaurantBusiness{store: store}
}

func (business *findRestaurantBusiness) FindRestaurant(context context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := business.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return result, nil
}