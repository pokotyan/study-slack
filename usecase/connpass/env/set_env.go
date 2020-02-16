package env

import (
	"context"
	"net/url"
	"os"
	"strconv"

	slackUtils "github.com/pokotyan/connpass-map-api/utils/slack"
)

type EnvError struct {
	Msg string
	ID  string
}

func SetEnv(ctx context.Context, rawBody string) []EnvError {
	message := slackUtils.ParseSubmissionCallBack(rawBody)
	webhookURL := slackUtils.GetSubmissionValue(message, "WEB_HOOK_URL")
	searchRange := slackUtils.GetSubmissionValue(message, "SEARCH_RANGE")
	numOfPeople := slackUtils.GetSubmissionValue(message, "NUM_OF_PEOPLE")

	Errors := []EnvError{}

	if _, err := url.ParseRequestURI(webhookURL); err != nil {
		Errors = append(Errors, EnvError{Msg: "有効なURLではありません。", ID: "WEB_HOOK_URL"})
	}

	if _, err := strconv.Atoi(searchRange); err != nil {
		Errors = append(Errors, EnvError{Msg: "数字で入力してください。", ID: "SEARCH_RANGE"})
	}

	if _, err := strconv.Atoi(numOfPeople); err != nil {
		Errors = append(Errors, EnvError{Msg: "数字で入力してください。", ID: "NUM_OF_PEOPLE"})
	}

	if len(Errors) != 0 {
		return Errors
	}

	os.Setenv("WEB_HOOK_URL", webhookURL)
	os.Setenv("SEARCH_RANGE", searchRange)
	os.Setenv("NUM_OF_PEOPLE", numOfPeople)

	return nil
}
