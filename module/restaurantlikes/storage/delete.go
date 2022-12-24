package restaurantlikestorage

import (
	restaurantlikesmodel "FoodDelivery/module/restaurantlikes/model"
	"context"
)

func (s *sqlStore) Delete(context context.Context, userId, restaurantId int) error {
	if err := s.db.Table(restaurantlikesmodel.Like{}.TableName()).Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		return err
	}
	return nil
}