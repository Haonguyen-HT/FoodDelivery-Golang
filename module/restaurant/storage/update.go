package restaurantstorage

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
	"fmt"
	"gorm.io/gorm"
)

func (s *sqlStore) Update(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	fmt.Println(*data)
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sqlStore) IncreaseLikeCount(context context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
func (s *sqlStore) DecreaseLikeCount(context context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}