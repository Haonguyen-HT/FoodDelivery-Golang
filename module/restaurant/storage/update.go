package retaurantstorage

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
	"fmt"
)

func (s *sqlStore) Update(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	fmt.Println(*data)
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}