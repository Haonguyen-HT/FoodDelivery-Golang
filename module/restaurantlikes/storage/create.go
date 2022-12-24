package restaurantlikestorage

import (
	"FoodDelivery/common"
	retauranlikesmodel "FoodDelivery/module/restaurantlikes/model"
	"context"
)

func (s *sqlStore) CreateLikeRestaurant(context context.Context, data *retauranlikesmodel.UserLike) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}