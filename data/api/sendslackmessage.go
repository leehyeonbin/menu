package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"menu-go/data/dto/slack"
)

func SendSlackMessage(token, channel, imageUrl string) error {
	url := "https://slack.com/api/chat.postMessage"

	blocks := []slack.Block {
		{
			Type: "image",
			ImageBlock: &slack.ImageBlock{
				Type: "image",
				ImageURL: imageUrl,
				AltText: "Image",
			},
		},
	}

	payload := slack.SlackMessage{
		Channel: channel,
		Blocks: blocks,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if error != nil {
		return fmt.Errorf("newRequest: %w", error)
	}

	request.Header.Set("Authorization", "Bearer " + token)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return fmt.Errorf("do: %w", error)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", response.StatusCode)
	}

	return nil
}