package models

type Table struct {
	Id        string             `json:"id"`
	HasHost   bool               `json:"has_host"`
	Players   map[string]*Player `json:"players"`
	WhoseTurn string             `json:"whose_turn"`
	Dealer    string             `json:"dealer"`
	Pot       int                `json:"pot"`
}

func newTable(id string) *Table {
	return &Table{
		Id:        id,
		HasHost:   false,
		Players:   make(map[string]*Player),
		WhoseTurn: "",
		Dealer:    "",
		Pot:       0,
	}
}
