package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"menu-go/data/dto/menu"
	"net/http"
)

func FetchMenuAPI(url string) (*menu.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

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
