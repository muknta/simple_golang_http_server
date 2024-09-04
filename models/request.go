package models

type Request struct {
	Cmd string `json:"cmd"`
	ID  int    `json:"id"`
}
