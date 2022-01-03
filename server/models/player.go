package models

type Player struct {
	Name   string `json:"name"`
	IsHost bool   `json:"is_host"`
	Chips  int    `json:"chips"`
	Spot   int    `json:"spot"`
}
