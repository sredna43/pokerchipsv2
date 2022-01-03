package models

type Table struct {
	Id        string             `json:"id"`
	HasHost   bool               `json:"has_host"`
	Players   map[string]*Player `json:"players"`
	WhoseTurn int                `json:"whose_turn"`
	Dealer    string             `json:"dealer"`
	Pot       int                `json:"pot"`
}

func NewTable(id string) *Table {
	return &Table{
		Id:        id,
		HasHost:   false,
		Players:   make(map[string]*Player),
		WhoseTurn: 0,
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
			Spot:   len(t.Players),
		}
		return true
	}
}

func (t *Table) RemovePlayer(name string) bool {
	player, ok := t.Players[name]
	if player.IsHost {
		t.HasHost = false
	}
	for _, p := range t.Players {
		if p.Spot > player.Spot {
			p.Spot -= 1
		}
		if !t.HasHost {
			p.IsHost = true
		}
	}
	delete(t.Players, name)
	return ok
}
