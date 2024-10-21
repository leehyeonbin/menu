package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"menu-go/dto"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	url := "https://hubkitchen.startup-plus.kr/api/cportal/board?size=10&page=0&portalSeqNo=395&boardId=B000100&title=%EB%A9%94%EB%89%B4"

	response, error := fetchAPI(url)
	if error != nil {
		log.Fatalf("fetchAPI: %v", error)
	}

	src, err := extractSrcFromHTML(response.Data.Contents[0].Contents)
	if err != nil {
		log.Fatalf("extractSrcFromHTML: %v", err)
	}

	downloadImage(src, "image.jpg")

	// VSCode에서 이미지 열기
	if err := openImageInVSCode("image.jpg"); err != nil {
		fmt.Println("Error opening image in VSCode:", err)
		return
	}
}

func extractSrcFromHTML(html string) (string, error) {
	// 정규 표현식으로 src 속성의 값만 추출
	re := regexp.MustCompile(`src="([^"]+)"`)
	matches := re.FindStringSubmatch(html)

	if len(matches) > 1 {
		return "https://hubkitchen.startup-plus.kr" + matches[1], nil
	} else {
		return "", fmt.Errorf("no src attribute found")
	}

}

func openImageInVSCode(filename string) error {
	// VSCode에서 이미지 열기 명령 실행
	cmd := exec.Command("code", filename)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to open image in VSCode: %w", err)
	}
	return nil
}

// 이미지 다운로드 함수
func downloadImage(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}

func fetchAPI(url string) (*dto.Response, error) {
	httpResponse, error := http.Get(url)

	if error != nil {
		return nil, fmt.Errorf("fetchAPI: %w", error)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", httpResponse.StatusCode)
	}

	body, error := io.ReadAll(httpResponse.Body)
	if error != nil {
		return nil, fmt.Errorf("error reading body: %v", error)
	}

	var response dto.Response
	if error := json.Unmarshal(body, &response); error != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", error)
	}

	return &response, nil
}
