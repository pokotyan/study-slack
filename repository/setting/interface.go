package repository

import (
	"context"

	"github.com/nlopes/slack"
	"github.com/pokotyan/study-slack/infrastructure/rdb/models"
)

type (
	SettingRepository interface {
		FetchCurrentSetting(ctx context.Context) models.SettingHistory
		Update(ctx context.Context, searchRange int, numOfPeople int, word string) models.SettingHistory
		MakeDialog(s models.SettingHistory, userID string) *slack.Dialog
		IsConfigured(s models.SettingHistory) bool
	}
)
