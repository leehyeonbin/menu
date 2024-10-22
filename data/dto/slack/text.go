package slack

// 텍스트 블록 정의 (선택적 사용)
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}