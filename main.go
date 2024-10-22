package main

import (
	"fmt"
	"log"
	"menu-go/data/api"
	"menu-go/util"
)

func main() {
	// slackToken := os.Getenv("SLACK_TOKEN")
	url := "https://hubkitchen.startup-plus.kr/api/cportal/board?size=10&page=0&portalSeqNo=395&boardId=B000100&title=%EB%A9%94%EB%89%B4"

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
	if err := util.OpenImageInVSCode("image.jpg"); err != nil {
		fmt.Println("Error opening image in VSCode:", err)
		return
	}
}
