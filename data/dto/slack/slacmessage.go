package slack

type SlackMessage struct {
	Channel string `json:"channel"`
	Blocks []Block `json:"blocks"`
}