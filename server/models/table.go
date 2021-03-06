package models

type Table struct {
	Id           string             `json:"id"`
	HasHost      bool               `json:"has_host"`
	Players      map[string]*Player `json:"players"`
	WhoseTurn    int                `json:"whose_turn"`
	Dealer       string             `json:"dealer"`
	DealerSpot   int                `json:"dealer_spot"`
	Pot          int                `json:"pot"`
	Playing      bool               `json:"playing"`
	InitialChips int                `json:"initial_chips"`
	CanCheck     bool               `json:"can_check"`
	BettingRound int                `json:"betting_round"`
	CurrentBet   int                `json:"current_bet"`
	HandWon      bool               `json:"hand_won"`
}

func NewTable(id string) *Table {
	return &Table{
		Id:           id,
		HasHost:      false,
		Players:      make(map[string]*Player),
		WhoseTurn:    0,
		Dealer:       "",
		DealerSpot:   0,
		Pot:          0,
		Playing:      false,
		InitialChips: 1000,
		CanCheck:     true,
		BettingRound: 0,
		HandWon:      false,
	}
}

func (t *Table) AddPlayer(name string) bool {
	make_host := len(t.Players) == 0
	if _, ok := t.Players[name]; ok {
		return false
	} else {
		if make_host {
			t.HasHost = true
			t.Dealer = name
		}
		t.Players[name] = &Player{
			Name:     name,
			IsHost:   make_host,
			IsDealer: make_host,
			Chips:    t.InitialChips,
			Spot:     len(t.Players),
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
		if player.IsDealer {
			t.DealerSpot = player.Spot
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
	t.InitialChips = amount
	return true
}

func (t *Table) SetDealer(newDealer string) bool {
	if p, ok := t.Players[t.Dealer]; ok {
		p.IsDealer = false
	}
	if p, ok := t.Players[newDealer]; ok {
		t.Dealer = newDealer
		p.IsDealer = true
		t.DealerSpot = p.Spot
		return true
	}
	return false
}

func (t *Table) StartGame() bool {
	if !t.Playing {
		t.WhoseTurn = 0
		t.Pot = 0
		t.Playing = true
		t.BettingRound = 0
		t.CanCheck = true
		t.CurrentBet = 0
		for _, p := range t.Players {
			p.Chips = t.InitialChips
			p.Folded = false
			if p.IsDealer {
				t.DealerSpot = p.Spot
			}
		}
		return true
	}
	return false
}

func (t *Table) NextTurn(raise bool) {
	t.WhoseTurn += 1
	if t.WhoseTurn == len(t.Players) {
		t.WhoseTurn = 0
	}
	for _, p := range t.Players {
		if p.Spot == t.WhoseTurn && p.Folded {
			t.NextTurn(false)
		}
	}
}

func (t *Table) Check(name string) bool {
	if !t.CanCheck || t.WhoseTurn != t.Players[name].Spot || t.HandWon {
		return false
	}
	t.NextTurn(false)
	return true
}

func (t *Table) Bet(name string, amount int) bool {
	p, ok := t.Players[name]
	if t.WhoseTurn != p.Spot || p.Chips < amount || !ok || t.HandWon {
		return false
	}
	t.CurrentBet = amount
	t.Pot += amount
	t.CanCheck = false
	p.Chips -= amount
	t.NextTurn(false)
	return true
}

func (t *Table) Call(name string) bool {
	p, ok := t.Players[name]
	if t.WhoseTurn != p.Spot || p.Chips < t.CurrentBet || !ok || t.HandWon {
		return false
	}
	t.Pot += t.CurrentBet
	t.CanCheck = false
	p.Chips -= t.CurrentBet
	t.NextTurn(false)
	return true
}

func (t *Table) Raise(name string, amount int) bool {
	p, ok := t.Players[name]
	if t.WhoseTurn != p.Spot || p.Chips < (t.CurrentBet+amount) || !ok || t.HandWon {
		return false
	}
	t.Pot += t.CurrentBet + amount
	t.CanCheck = false
	p.Chips -= (t.CurrentBet + amount)
	t.NextTurn(true)
	return true
}

func (t *Table) Fold(name string) bool {
	p, ok := t.Players[name]
	if t.WhoseTurn != p.Spot || p.Folded || !ok || t.HandWon {
		return false
	}
	p.Folded = true
	t.NextTurn(false)
	return true
}

func (t *Table) WinRound(name string) bool {
	p, ok := t.Players[name]
	if !ok {
		return false
	}
	t.HandWon = true
	p.Chips += t.Pot
	return true
}

func (t *Table) NewRound() bool {
	var nextDealer string
	t.DealerSpot += 1
	newDealerSpot := t.DealerSpot
	if newDealerSpot == len(t.Players) {
		newDealerSpot = 0
	}
	for _, p := range t.Players {
		p.Folded = false
		if p.Spot == newDealerSpot {
			nextDealer = p.Name
		}
	}
	t.BettingRound = 0
	t.CanCheck = true
	t.CurrentBet = 0
	t.SetDealer(nextDealer)
	t.Pot = 0
	t.HandWon = false
	t.WhoseTurn = t.Players[t.Dealer].Spot + 1
	if t.WhoseTurn == len(t.Players) {
		t.WhoseTurn = 0
	}
	return true
}
