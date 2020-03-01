package repository

import (
	"context"
	"time"

	"github.com/pokotyan/study-slack/infrastructure/rdb/models"
	cc "github.com/pokotyan/study-slack/utils/context"
)

type (
	settingRepository struct{}
)

func NewSettingRepository() SettingRepository {
	return &settingRepository{}
}

func (*settingRepository) FetchCurrentSetting(ctx context.Context) models.SettingHistory {
	db, _ := cc.GetDB(ctx)

	setting := models.SettingHistory{}
	db.Last(&setting)

	return setting
}

func (*settingRepository) Update(ctx context.Context, searchRange int, numOfPeople int) models.SettingHistory {
	db, _ := cc.GetDB(ctx)

	setting := models.SettingHistory{}
	setting.SearchRange = searchRange
	setting.NumOfPeople = numOfPeople
	setting.CreatedAt = time.Now()
	setting.UpdatedAt = time.Now()

	db.Create(&setting)

	return setting
}
