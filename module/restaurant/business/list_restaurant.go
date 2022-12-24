package restaurantbusiness

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBusiness(repo ListRestaurantRepo) *listRestaurantBusiness {
	return &listRestaurantBusiness{repo: repo}
}

func (business *listRestaurantBusiness) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := business.repo.ListRestaurant(context, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}