package api

import (
	"fmt"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func SendSlackMessage(Token, channel, titile string, imagePath string) error {
	api := slack.New(Token)

	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	defer file.Close()

	// 파일 정보를 얻기 위해 파일 사이즈 계산
	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info: %v", err)
	}

	params := slack.UploadFileV2Parameters{
		Filename: "menu",
		Channel:  channel,
		Title:    titile,
		FileSize: int(stat.Size()),
		File:     imagePath,
	}

	FileSummary, err := api.UploadFileV2(params)
	if err != nil {
		fmt.Printf("api Error %s\n", err)
		return err
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("Image %s successfully sent to channel %s at %s", FileSummary.Title, channel, timestamp)
	return nil
}
