package userstore

import (
	"FoodDelivery/common"
	usermodel "FoodDelivery/module/user/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(err)
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}