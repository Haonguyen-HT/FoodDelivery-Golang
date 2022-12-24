package userstore

import (
	"FoodDelivery/common"
	usermodel "FoodDelivery/module/user/model"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}