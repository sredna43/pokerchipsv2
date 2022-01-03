package main

import (
	"encoding/json"
	"fmt"

	"github.com/sredna43/pokerchipsv2/models"
)

type Request struct {
	PlayerName string `json:"name"`
	Action     string `json:"action"`
	Amount     int    `json:"amount"`
}

type Response struct {
	Requester string                    `json:"requester"`
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
	res.Requester = req.PlayerName

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
			res.Message = req.PlayerName + " left"
		}
	case "set_dealer":
		t.Dealer = req.PlayerName
		res.Message = req.PlayerName + " is now dealer"
	case "set_spot":
		if player, ok := t.Players[req.PlayerName]; ok {
			player.Spot = req.Amount
			res.Message = fmt.Sprintf("Set %s's spot to %d", req.PlayerName, req.Amount)
		}
	case "start_game":
		res.Message = "START_GAME"
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
