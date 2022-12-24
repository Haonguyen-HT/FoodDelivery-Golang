package restaurantrepo

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantRepo struct {
	store ListRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store}
}

func (business *listRestaurantRepo) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := business.store.ListDataWithCondition(context, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	//var ids = make([]int, len(result))
	//
	//for i := range ids {
	//	ids[i] = result[i].Id
	//}
	//
	//likeMap, err := business.likeStore.GetRestaurantLikes(context, ids)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return result, nil
	//}
	//
	//for i, item := range result {
	//	result[i].LikedCount = likeMap[item.Id]
	//}

	return result, nil
}