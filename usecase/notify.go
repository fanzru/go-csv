package usecase

import (
	"fmt"
	"go-csv/domain/models"
	"go-csv/infrastructure/config"

	"github.com/nlopes/slack"
)

func DoNotifySlackSend(param models.NotifSlackParam) error {

	attachment := slack.Attachment{
		Text: param.Text,
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	slackEnv, err := config.GetSlackEnv()
	if err != nil {
		return err
	}
	fmt.Println("-------------------------------- env", slackEnv.SlackWebhook)
	err = slack.PostWebhook(slackEnv.SlackWebhook, &msg)
	if err != nil {
		return err
	}

	return nil
}
