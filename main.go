package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://hubkitchen.startup-plus.kr/api/cportal/board?size=10&page=0&portalSeqNo=395&boardId=B000100&title=%EB%A9%94%EB%89%B4"

	response, error := http.Get(url)
	if error != nil {
		log.Fatalf(" Error Occurred. %v", error)
	}
	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)
	if error != nil {
		log.Fatalf("Failed to read response: %v", error)
	}

	fmt.Println(string(body))
}
