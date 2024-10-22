package util

import (
	"fmt"
	"regexp"
)

func ExtractSrcFromHTML(html string) (string, error) {
	// 정규 표현식으로 src 속성의 값만 추출
	re := regexp.MustCompile(`src="([^"]+)"`)
	matches := re.FindStringSubmatch(html)

	if len(matches) > 1 {
		return "https://hubkitchen.startup-plus.kr" + matches[1], nil
	} else {
		return "", fmt.Errorf("no src attribute found")
	}

}
