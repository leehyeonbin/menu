package dto

type Response struct {
	Result bool `json:"result"`
	Data   Data `json:"data"`
}
