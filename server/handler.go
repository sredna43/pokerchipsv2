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
	Requester string        `json:"requester"`
	Message   string        `json:"message"`
	Error     bool          `json:"error"`
	Table     *models.Table `json:"table"`
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
		if t.AddPlayer(req.PlayerName) {
			res.Message = fmt.Sprintf("%s joined", req.PlayerName)
		} else {
			res.Message = fmt.Sprintf("Player %s already exists", req.PlayerName)
			res.Error = true
		}
	case "remove_player":
		if t.RemovePlayer(req.PlayerName) {
			res.Message = fmt.Sprintf("%s left", req.PlayerName)
		} else {
			res.Message = fmt.Sprintf("Player %s doesn't exist", req.PlayerName)
			res.Error = true
		}
	case "set_dealer":
		if t.SetDealer(req.PlayerName) {
			res.Message = fmt.Sprintf("%s is now dealer", req.PlayerName)
		} else {
			res.Message = fmt.Sprintf("Could not make %s dealer", req.PlayerName)
			res.Error = true
		}
	case "move_player":
		if t.MovePlayer(req.PlayerName, req.Amount) {
			res.Message = fmt.Sprintf("Moved player %s", req.PlayerName)
		} else {
			res.Message = fmt.Sprintf("Could not move player %s", req.PlayerName)
			res.Error = true
		}
	case "set_initial_chips":
		if t.SetInitialChips(req.Amount) {
			res.Message = fmt.Sprintf("Set initial chips to %d", req.Amount)
		} else {
			res.Message = fmt.Sprintf("Could not set initial chips to %d", req.Amount)
			res.Error = true
		}
	case "start_game":
		if t.StartGame() {
			res.Message = "START_GAME"
		} else {
			res.Message = "Could not start game, already playing"
			res.Error = true
		}
	case "check":
		if t.Check(req.PlayerName) {
			res.Message = fmt.Sprintf("%s checks", req.PlayerName)
		} else {
			res.Message = "Could not check"
			res.Error = true
		}
	case "bet":
		if t.Bet(req.PlayerName, req.Amount) {
			res.Message = fmt.Sprintf("%s bets %d", req.PlayerName, req.Amount)
		} else {
			res.Message = "Could not bet"
			res.Error = true
		}
	case "call":
		if t.Call(req.PlayerName) {
			res.Message = fmt.Sprintf("%s calls", req.PlayerName)
		} else {
			res.Message = "Could not call"
			res.Error = true
		}
	case "raise":
		if t.Raise(req.PlayerName, req.Amount) {
			res.Message = fmt.Sprintf("%s raises %d", req.PlayerName, req.Amount)
		} else {
			res.Message = "Could not raise"
			res.Error = true
		}
	case "fold":
		if t.Fold(req.PlayerName) {
			res.Message = fmt.Sprintf("%s folds", req.PlayerName)
		} else {
			res.Message = "Could not fold"
			res.Error = true
		}
	case "win_round":
		if t.WinRound(req.PlayerName) {
			res.Message = fmt.Sprintf("%s wins the round", req.PlayerName)
		} else {
			res.Message = "Could not choose winner"
			res.Error = true
		}
	case "new_round":
		if t.NewRound() {
			res.Message = "NEW_ROUND"
		} else {
			res.Message = "Could not start new round"
			res.Error = true
		}
	case "game_state":
		res.Message = "current gamestate"
	default:
		res.Message = "Unknown action: " + req.Action + " :: " + string(r)
		res.Error = true
	}

	if !res.Error {
		res.Table = t
	}

	if b, err := json.Marshal(res); err != nil {
		return []byte("Error building response")
	} else {
		return b
	}
}
