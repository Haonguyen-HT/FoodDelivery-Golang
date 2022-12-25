package restaurantmodel

import (
	"FoodDelivery/common"
	"errors"
	"strings"
)

type Restaurant struct {
	common.SQLModel
	Name       string             `json:"name" gorm:"column:name;" `
	Addr       string             `json:"addr" gorm:"column:addr;"`
	Logo       *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover      *common.Images     `json:"cover" gorm:"column:cover;"`
	UserID     int                `json:"-" gorm:"column:user_id;"`
	User       *common.SimpleUser `json:"owner" gorm:"preload:false; "`
	LikedCount int                `json:"liked_count" gorm:"column:liked_count;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)

	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;" `
	Addr            string         `json:"addr" gorm:"column:addr;"`
	UserID          int            `json:"-" gorm:"column:user_id;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

const EntityName = "Restaurant"

func (data RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name    *string        `json:"name" gorm:"column:name;"`
	Address *string        `json:"address" gorm:"column:addr;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var ErrNameIsEmpty = errors.New("name is not empty")