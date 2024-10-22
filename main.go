package main

import (
	"log"
	"os"
	"menu-go/data/api"
	"github.com/joho/godotenv"
	"menu-go/util"
)

func main() {
	    // .env 파일 로드
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
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

	// VSCode에서 이미지 열기
	// if err := util.OpenImageInVSCode("image.jpg"); err != nil {
	// 	fmt.Println("Error opening image in VSCode:", err)
	// 	return
	// }

	api.SendSlackMessage(slackToken, channelId, src)
}
