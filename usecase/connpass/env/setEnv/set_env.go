package env

import (
	"context"
	"strconv"

	settingRepo "github.com/pokotyan/study-slack/repository/setting"
	slackUtils "github.com/pokotyan/study-slack/utils/slack"
)

type EnvError struct {
	Msg string
	ID  string
}

func NewSetEnvImpl(sr settingRepo.SettingRepository) ConnpassEnvUsecase {
	return &connpassEnvUsecaseImpl{
		settingRepo: sr,
	}
}

func (c *connpassEnvUsecaseImpl) SetEnv(ctx context.Context, rawBody string) []EnvError {
	message := slackUtils.ParseSubmissionCallBack(rawBody)
	searchRange := slackUtils.GetSubmissionValue(message, "SEARCH_RANGE")
	numOfPeople := slackUtils.GetSubmissionValue(message, "NUM_OF_PEOPLE")

	Errors := []EnvError{}

	sr, err := strconv.Atoi(searchRange)
	if err != nil {
		Errors = append(Errors, EnvError{Msg: "数字で入力してください。", ID: "SEARCH_RANGE"})
	}

	if sr > 30 {
		Errors = append(Errors, EnvError{Msg: "MAXは30日です。", ID: "SEARCH_RANGE"})
	}

	nop, err := strconv.Atoi(numOfPeople)
	if err != nil {
		Errors = append(Errors, EnvError{Msg: "数字で入力してください。", ID: "NUM_OF_PEOPLE"})
	}

	if len(Errors) != 0 {
		return Errors
	}

	c.settingRepo.Update(ctx, sr, nop)

	return nil
}
