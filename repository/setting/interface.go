package repository

import (
	"context"

	"github.com/pokotyan/study-slack/infrastructure/rdb/models"
)

type (
	SettingRepository interface {
		FetchCurrentSetting(ctx context.Context) models.SettingHistory
		Update(ctx context.Context, searchRange int, numOfPeople int) models.SettingHistory
	}
)
