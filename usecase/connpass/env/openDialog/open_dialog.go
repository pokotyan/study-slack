package env

import (
	"context"
	"fmt"

	"github.com/nlopes/slack"
	settingRepo "github.com/pokotyan/study-slack/repository/setting"
)

func NewOpenDialogImpl(client *slack.Client, sr settingRepo.SettingRepository) ConnpassEnvUsecase {
	return &connpassEnvUsecaseImpl{
		slackClient: client,
		settingRepo: sr,
	}
}

func (u connpassEnvUsecaseImpl) OpenDialog(ctx context.Context, userID string, triggerID string) error {
	s := u.settingRepo.FetchCurrentSetting(ctx)
	dialog := u.settingRepo.MakeDialog(s, userID)

	if err := u.slackClient.OpenDialog(triggerID, *dialog); err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}
