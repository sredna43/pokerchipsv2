package models

type Table struct {
	Id        string             `json:"id"`
	HasHost   bool               `json:"has_host"`
	Players   map[string]*Player `json:"players"`
	WhoseTurn int                `json:"whose_turn"`
	Dealer    string             `json:"dealer"`
	Pot       int                `json:"pot"`
	Playing   bool               `json:"playing"`
}

func NewTable(id string) *Table {
	return &Table{
		Id:        id,
		HasHost:   false,
		Players:   make(map[string]*Player),
		WhoseTurn: 0,
		Dealer:    "",
		Pot:       0,
		Playing:   false,
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

func (t *Table) MovePlayer(name string, direction int) bool {
	if player, ok := t.Players[name]; ok {
		if player.Spot+direction < 0 || player.Spot+direction >= len(t.Players) {
			return false
		}
		oldSpot := player.Spot
		newSpot := player.Spot + direction
		for _, p := range t.Players {
			if p.Spot == newSpot {
				p.Spot = oldSpot
			}
			if p.Name == player.Name {
				p.Spot = newSpot
			}
		}
		return true
	}
	return false
}

func (t *Table) SetInitialChips(amount int) bool {
	if t.Playing || amount < 0 {
		return false
	}
	for _, p := range t.Players {
		p.Chips = amount
	}
	return true
}

func (t *Table) SetDealer(newDealer string) bool {
	if p, ok := t.Players[t.Dealer]; ok {
		p.IsDealer = false
	}
	if p, ok := t.Players[newDealer]; ok {
		t.Dealer = newDealer
		p.IsDealer = true
		return true
	}
	return false
}

func (t *Table) StartGame() bool {
	if !t.Playing {
		t.WhoseTurn = 0
		t.Pot = 0
		t.Playing = true
		return true
	} else {
		return false
	}
}
