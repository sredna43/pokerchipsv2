package main

import (
	"encoding/json"

	"github.com/sredna43/pokerchipsv2/models"
)

type Request struct {
	PlayerName string `json:"name"`
	Action     string `json:"action"`
	Amount     int    `json:"amount"`
}

type Response struct {
	Players   map[string]*models.Player `json:"players"`
	Message   string                    `json:"message"`
	Pot       int                       `json:"pot"`
	WhoseTurn string                    `json:"whose_turn"`
	Dealer    string                    `json:"dealer"`
	Error     bool                      `json:"error"`
}

func handleRequest(t *models.Table, r []byte) []byte {
	req := &Request{}
	res := &Response{Error: false}
	if err := json.Unmarshal(r, req); err != nil {
		res.Message = "Error: " + err.Error()
	}

	switch req.Action {
	case "add_player":
		if !t.AddPlayer(req.PlayerName) {
			res.Message = "Player " + req.PlayerName + " already exists"
			res.Error = true
		} else {
			res.Message = req.PlayerName + " joined"
		}
	case "remove_player":
		if !t.RemovePlayer(req.PlayerName) {
			res.Message = "Player " + req.PlayerName + " doesn't exist"
			res.Error = true
		} else {
			res.Message = req.PlayerName + " removed"
		}
	case "game_state":
		res.Players = t.Players
		res.Pot = t.Pot
		res.WhoseTurn = t.WhoseTurn
		res.Dealer = t.Dealer
		res.Message = "current gamestate"
	default:
		res.Message = "Unknown action: " + req.Action + " :: " + string(r)
		res.Error = true
	}

	res.Players = t.Players

	if b, err := json.Marshal(res); err != nil {
		return []byte("Error building response")
	} else {
		return b
	}
}
