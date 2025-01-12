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

	response, error := api.FetchMenuAPI(url)
	if error != nil {
		log.Fatalf("fetchAPI: %v", error)
	}

	src, err := util.ExtractSrcFromHTML(response.Data.Contents[0].Contents)
	if err != nil {
		log.Fatalf("extractSrcFromHTML: %v", err)
	}

	util.DownloadImage(src, "image.jpg")

	api.SendSlackMessage(slackToken, channelId, response.Data.Contents[0].Title, "image.jpg")
}
