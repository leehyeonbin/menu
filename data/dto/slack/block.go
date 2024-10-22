package slack

type Block struct {
	Type   string `json:"type"`
	Text *Text  `json:"text,omitempty"`
	ImageBlock *ImageBlock `json:"image,omitempty"`

}