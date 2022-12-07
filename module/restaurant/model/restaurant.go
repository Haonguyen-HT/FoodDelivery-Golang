package restaurantmodel

import (
	"FoodDelivery/common"
	"errors"
	"strings"
)

type Restaurant struct {
	common.SQLModel
	Name string `json:"name" gorm:"column:name;" `
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;" `
	Addr            string `json:"addr" gorm:"column:addr;"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

const EntityName = "Restaurant"

func (data RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return errors.New("name is not empty")
	}

	return nil
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}