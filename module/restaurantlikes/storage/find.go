package restaurantlikestorage

import (
	restaurantlikesmodel "FoodDelivery/module/restaurantlikes/model"
	"context"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantlikesmodel.Like, error) {
	var data restaurantlikesmodel.Like

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}