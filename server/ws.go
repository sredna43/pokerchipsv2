package main

import (
	"log"

	"github.com/sredna43/pokerchipsv2/models"
)

type Ws struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	table      *models.Table
	quit       chan bool
}

func NewServer(tableId string) *Ws {
	return &Ws{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		table:      models.NewTable(tableId),
		quit:       make(chan bool),
	}
}

func (ws *Ws) run() {
	for {
		select {
		case client := <-ws.register:
			log.Println("Register client")
			ws.clients[client] = true
		case client := <-ws.unregister:
			log.Println("Unregister client")
			if _, ok := ws.clients[client]; ok {
				delete(ws.clients, client)
				close(client.send)
			}
			if len(ws.clients) == 0 {
				if ws.table.Id == "test" {
					continue
				}
				log.Println("Quitting due to no more players in lobby")
				delete(Servers, ws.table.Id)
				return
			}
		case request := <-ws.broadcast:
			response := handleRequest(ws.table, request)
			for client := range ws.clients {
				select {
				case client.send <- response:
				default:
					close(client.send)
					delete(ws.clients, client)
				}
			}
		case <-ws.quit:
			log.Println("Quitting ws due to request")
			delete(Servers, ws.table.Id)
			return
		}
	}
}
