package slackUtils

import (
	"encoding/json"
	"fmt"
	"strings"

	slackWebHook "github.com/ashwanthkumar/slack-go-webhook"
	slack "github.com/nlopes/slack"
)

type Slack struct {
	WebhookURL string
}

func NewSlack(webhookURL string) (*Slack, error) {
	return &Slack{
		WebhookURL: webhookURL,
	}, nil
}

func ParseSubmissionCallBack(str string) slack.InteractionCallback {
	str = strings.Replace(str, "payload=", "", 1)

	var message slack.InteractionCallback

	json.Unmarshal([]byte(str), &message)

	return message
}

func GetSubmissionValue(message slack.InteractionCallback, id string) string {
	return message.DialogSubmissionCallback.Submission[id]
}

func (s Slack) SendToSlack(text string, sendable bool, setAttach func(attachment *slackWebHook.Attachment)) error {
	attachment := slackWebHook.Attachment{}

	setAttach(&attachment)

	payload := slackWebHook.Payload{
		Text:        text,
		IconEmoji:   ":monkey_face:",
		Attachments: []slackWebHook.Attachment{attachment},
	}

	if sendable {
		err := slackWebHook.Send(s.WebhookURL, "", payload)
		if len(err) > 0 {
			fmt.Printf("error: %s\n", err)
		}
	}

	return nil
}
