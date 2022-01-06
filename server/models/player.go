package models

type Player struct {
	Name     string `json:"name"`
	IsHost   bool   `json:"is_host"`
	IsDealer bool   `json:"is_dealer"`
	Chips    int    `json:"chips"`
	Spot     int    `json:"spot"`
	Folded   bool   `json:"folded"`
}
