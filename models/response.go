package models

type Response struct {
	Cmd    string `json:"cmd"`
	ID     int    `json:"id"`
	Result string `json:"result"`
}
