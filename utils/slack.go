package utils

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
)

type Slack struct {
	WebhookURL string
}

func NewSlack(webhookURL string) (*Slack, error) {
	return &Slack{
		WebhookURL: webhookURL,
	}, nil
}

func (s Slack) SendToSlack(text string, sendable bool) error {
	attachment1 := slack.Attachment{}
	// attachment1.AddField(slack.Field{Title: "Author", Value: "Ashwanth Kumar"}).AddField(slack.Field{Title: "Status", Value: "Completed"})
	// attachment1.AddAction(slack.Action{Type: "button", Text: "Book flights 🛫", Url: "https://flights.example.com/book/r123456", Style: "primary"})
	// attachment1.AddAction(slack.Action{Type: "button", Text: "Cancel", Url: "https://flights.example.com/abandon/r123456", Style: "danger"})
	payload := slack.Payload{
		Text:        text,
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment1},
	}

	if sendable {
		err := slack.Send(s.WebhookURL, "", payload)
		if len(err) > 0 {
			fmt.Printf("error: %s\n", err)
		}
	}

	return nil
}
