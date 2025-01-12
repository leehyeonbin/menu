package menu

import (
	"log"
	"menu-go/feature/menu/data/remote/api"
	"menu-go/util"
	"os"
)

func Menu() {
	// .env 파일 로드
	slackToken := os.Getenv("SLACK_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")
	url := os.Getenv("URL")

	response, remoteError := api.FetchMenuAPI(url)
	if remoteError != nil {
		log.Fatalf("fetchAPI: %v", remoteError)
	}

	src, err := util.ExtractSrcFromHTML(response.Data.Contents[0].Contents)
	if err != nil {
		log.Fatalf("extractSrcFromHTML: %v", err)
	}

	imageDownError := util.DownloadImage(src, "image.jpg")
	if imageDownError != nil {
		log.Fatalf("extractSrcFromHTML: %v", err)
		return
	}

	sendError := api.SendSlackMessage(slackToken, channelId, response.Data.Contents[0].Title, "image.jpg")
	if sendError != nil {
		log.Fatalf("extractSrcFromHTML: %v", err)
		return
	}
}
