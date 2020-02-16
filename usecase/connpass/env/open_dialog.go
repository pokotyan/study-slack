package env

import (
	"context"
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func NewOpenDialogImpl(client *slack.Client) ConnpassEnvUsecase {
	return &connpassEnvUsecaseImpl{
		slackClient: client,
	}
}

func makeDialog(userID string) *slack.Dialog {
	return &slack.Dialog{
		Title:       "環境設定",
		SubmitLabel: "Submit",
		CallbackID:  userID + "EnvForm",
		Elements: []slack.DialogElement{slack.DialogInput{
			Label:       "WEB_HOOK_URL",
			Type:        slack.InputTypeText,
			Name:        "WEB_HOOK_URL",
			Placeholder: os.Getenv("WEB_HOOK_URL"),
			Hint:        "通知したいチャンネルのwebHookURL",
		},
			slack.DialogInput{
				Label:       "検索範囲",
				Type:        slack.InputTypeText,
				Name:        "SEARCH_RANGE",
				Placeholder: os.Getenv("SEARCH_RANGE"),
				Hint:        "数字。どのくらい先までの勉強会を通知するか。7にすると１週間先までの勉強会を通知",
			},
			slack.DialogInput{
				Label:       "最低参加人数",
				Type:        slack.InputTypeText,
				Name:        "NUM_OF_PEOPLE",
				Placeholder: os.Getenv("NUM_OF_PEOPLE"),
				Hint:        "数字。100にすると参加人数が100人以上の勉強会のみ通知",
			},
		},
	}
}

func (u connpassEnvUsecaseImpl) OpenDialog(ctx context.Context, userID string, triggerID string) error {
	dialog := makeDialog(userID)

	if err := u.slackClient.OpenDialog(triggerID, *dialog); err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}
