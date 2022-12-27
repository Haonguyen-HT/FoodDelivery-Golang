package restaurantlikesmodel

import (
	"FoodDelivery/common"
	"fmt"
	"time"
)

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreateAt     *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

type UserLike struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int `json:"user_id" gorm:"column:user_id;"`
}

func (UserLike) TableName() string {
	return Like{}.TableName()
}

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func (l *UserLike) GetRestaurantId() int {
	return l.RestaurantId
}
func (l *UserLike) GetUserId() int {
	return l.UserId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"))
}
func ErrCannotDisLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("cannot unlike this restaurant"),
		fmt.Sprintf("ErrCannotUnLikeRestaurant"))
}