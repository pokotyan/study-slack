package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/nlopes/slack"
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

func (*settingRepository) Update(ctx context.Context, searchRange int, numOfPeople int, word string) models.SettingHistory {
	db, _ := cc.GetDB(ctx)

	setting := models.SettingHistory{}
	setting.SearchRange = searchRange
	setting.NumOfPeople = numOfPeople
	setting.Word = word
	setting.CreatedAt = time.Now()
	setting.UpdatedAt = time.Now()

	db.Create(&setting)

	return setting
}

func (*settingRepository) IsConfigured(s models.SettingHistory) bool {
	return s.SearchRange != 0 && s.NumOfPeople != 0
}

func (*settingRepository) MakeDialog(s models.SettingHistory, userID string) *slack.Dialog {
	return &slack.Dialog{
		Title:       "環境設定",
		SubmitLabel: "Submit",
		CallbackID:  userID + "EnvForm",
		Elements: []slack.DialogElement{
			slack.DialogInput{
				Label:       "検索範囲（現在の値: " + strconv.Itoa(s.SearchRange) + ")",
				Type:        slack.InputTypeText,
				Name:        "SEARCH_RANGE",
				Placeholder: strconv.Itoa(s.SearchRange),
				Hint:        "数字（日数）。どのくらい先までの勉強会を通知するか。ex) 7にすると１週間先までの勉強会を通知。",
			},
			slack.DialogInput{
				Label:       "最低参加人数（現在の値: " + strconv.Itoa(s.NumOfPeople) + ")",
				Type:        slack.InputTypeText,
				Name:        "NUM_OF_PEOPLE",
				Placeholder: strconv.Itoa(s.NumOfPeople),
				Hint:        "数字（人数）。ex) 100にすると参加人数が100人以上の勉強会のみ通知。",
			},
			slack.DialogInput{
				Label:       "検索ワード（現在の値: " + s.Word + ")",
				Type:        slack.InputTypeText,
				Name:        "WORD",
				Placeholder: s.Word,
				Hint:        "検索ワード（省略可）",
				Optional:    true,
			},
		},
	}
}
