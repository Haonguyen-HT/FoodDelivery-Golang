package restaurantmodel

import (
	"testing"
)

func TestRestaurantCreate_Validate(t *testing.T) {
	dataTest := RestaurantCreate{
		Name: "",
	}

	err := dataTest.Validate()

	if err == ErrNameIsEmpty {
		t.Error("Validate restaurant. Input name:", dataTest.Name, ". Expect: ErrNameIsEmpty", "Output: ", err)
		return
	}

	t.Log("Validate restaurant: passed")

}