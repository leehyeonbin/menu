package api

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"menu-go/data/dto/menu"
	"net/http"
)

func FetchMenuAPI(url string) (*menu.Response, error) {
	// 시스템의 루트 인증서 풀 로드
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// 사용자 정의 Transport 생성
	config := &tls.Config{
		RootCAs: rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	// GET 요청 수행
	httpResponse, error := client.Get(url)
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

	var response menu.Response
	if error := json.Unmarshal(body, &response); error != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", error)
	}

	return &response, nil
}
