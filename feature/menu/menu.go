package menu

import (
	"fmt"
	"log"
	"menu-go/feature/menu/data/remote/api"
	"menu-go/util"
	"os"
)

func Menu() {
	// .env 파일 로드
	fmt.Println("load Env file")
	slackToken := os.Getenv("SLACK_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")
	url := os.Getenv("URL")

	fmt.Println("fetch menu api")
	response, remoteError := api.FetchMenuAPI(url)
	if remoteError != nil {
		log.Fatalf("fetch menu api error: %v", remoteError)
	}

	fmt.Println("extract src from html")
	src, err := util.ExtractSrcFromHTML(response.Data.Contents[0].Contents)
	if err != nil {
		log.Fatalf("extract src from html error: %v", err)
	}

	fmt.Println("download image")
	imageDownError := util.DownloadImage(src, "image.jpg")
	if imageDownError != nil {
		log.Fatalf("download image error: %v", err)
		return
	}

	fmt.Println("send slack message")
	sendError := api.SendSlackMessage(slackToken, channelId, response.Data.Contents[0].Title, "image.jpg")
	if sendError != nil {
		log.Fatalf("send slack message error: %v", err)
		return
	}

	fmt.Println("Success Menu Process")
}
