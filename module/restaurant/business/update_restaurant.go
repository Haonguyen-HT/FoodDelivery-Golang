package restaurantbusiness

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
	"errors"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Update(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBusiness struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBusiness(store UpdateRestaurantStore) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{store: store}
}

func (business *updateRestaurantBusiness) UpdateRestaurant(
	context context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	oldData, err := business.store.FindDataWithCondition(
		context,
		map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}
	if err := business.store.Update(context, id, data); err != nil {
		return err
	}

	return nil
}