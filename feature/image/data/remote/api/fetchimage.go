package api

import (
	"fmt"
	"io"
	"net/http"
)

func FetchImage(url string) (readCloser io.ReadCloser, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %w", err)
	}
	return resp.Body, nil
}
