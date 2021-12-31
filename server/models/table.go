package models

type Table struct {
	Id        string             `json:"id"`
	HasHost   bool               `json:"has_host"`
	Players   map[string]*Player `json:"players"`
	WhoseTurn string             `json:"whose_turn"`
	Dealer    string             `json:"dealer"`
	Pot       int                `json:"pot"`
}

func NewTable(id string) *Table {
	return &Table{
		Id:        id,
		HasHost:   false,
		Players:   make(map[string]*Player),
		WhoseTurn: "",
		Dealer:    "",
		Pot:       0,
	}
}

func (t *Table) AddPlayer(name string) bool {
	make_host := len(t.Players) == 0
	if _, ok := t.Players[name]; ok {
		return false
	} else {
		t.Players[name] = &Player{
			Name:   name,
			IsHost: make_host,
		}
		return true
	}
}

func (t *Table) RemovePlayer(name string) bool {
	_, ok := t.Players[name]
	delete(t.Players, name)
	return ok
}
