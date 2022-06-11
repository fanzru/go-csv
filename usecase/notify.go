package usecase

import (
	"encoding/csv"
	"fmt"
	"go-csv/domain/models"
	"go-csv/infrastructure/config"
	"os"

	// "github.com/nlopes/slack"
	"github.com/slack-go/slack"
)

func DoNotifySlackSend(param models.NotifSlackParam) error {

	// ----------------------------------------------------------------
	csvFile, err := os.Create("./data.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	for i, usance := range param.Data {
		if i == 0 {
			var title []string
			title = append(title, "nama")
			title = append(title, "nim")
			title = append(title, "kelas")
			writer.Write(title)
		}
		var row []string
		row = append(row, usance.Nama)
		row = append(row, usance.NIM)
		row = append(row, usance.Kelas)
		writer.Write(row)
	}

	// remember to flush!
	writer.Flush()

	// --------------------------------- Text Attachment-------------------------------

	// attachment := slack.Attachment{
	// 	Text:  param.Text,
	// 	Title: "Halooooooo",
	// }
	// msg := slack.WebhookMessage{
	// 	Attachments: []slack.Attachment{attachment},
	// }

	// slackEnv, err := config.GetSlackEnv()
	// if err != nil {
	// 	return err
	// }

	// err = slack.PostWebhook(slackEnv.SlackWebhook, &msg)
	// if err != nil {
	// 	return err
	// }
	// --------------------------------- File Attachment-------------------------------
	slackEnv, err := config.GetSlackEnv()
	if err != nil {
		return err
	}
	api := slack.New(slackEnv.BotToken)
	ch := []string{slackEnv.ChId}

	params := slack.FileUploadParameters{
		Channels:       ch,
		File:           "./data.csv",
		Filetype:       "csv",
		InitialComment: "Halo guys punten mau kirim file....",
	}
	_, err = api.UploadFile(params)
	if err != nil {
		return err
	}

	return nil
}
