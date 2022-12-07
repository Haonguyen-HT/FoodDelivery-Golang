package retaurantstorage

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}