package restaurantmodel

type Filter struct {
	Status int `json:"status,omitempty" form:"status"`
}