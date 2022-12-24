package restaurantstorage

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var data []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1,2,3,4,5,6)")

	if f := filter; f != nil {
		if f.Status > 0 {
			db = db.Where("status = ?", f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.Cursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, err
	}

	if len(data) > 0 {
		last := data[len(data)-1]
		last.Mask(false)
		paging.NextCursor = last.UID.String()
	}

	return data, nil
}