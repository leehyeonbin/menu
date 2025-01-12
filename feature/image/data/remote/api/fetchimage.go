package api

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func FetchImage(url string) (readCloser io.ReadCloser, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %w", err)
	}
	return resp.Body, nil
}
