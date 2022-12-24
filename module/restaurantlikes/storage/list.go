package restaurantlikestorage

import (
	"FoodDelivery/common"
	restaurantlikesmodel "FoodDelivery/module/restaurantlikes/model"
	"context"
)

func (s *sqlStore) GetRestaurantLikes(
	context context.Context,
	ids []int,
) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		LikeCount    int `gorm:"column:count;"`
	}

	var listLike []sqlData

	if err := s.db.Table(restaurantlikesmodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil

}