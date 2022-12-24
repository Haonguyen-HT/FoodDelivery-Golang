package usermodel

import (
	"FoodDelivery/common"
	"errors"
)

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already exits"),
		"email has already exits",
		"ErrEmailExist")

	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid")
)