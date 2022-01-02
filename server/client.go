package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws   *Ws
	conn *websocket.Conn
	send chan []byte
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
	allowedOrigin  = "http://localhost:3000"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		log.Println(origin)
		return origin == allowedOrigin || origin == "http://localhost:3000"
	},
}

// Read from socket
func (c *Client) readPump() {
	defer func() {
		c.ws.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Socket closed: %v", err)
			}
			break
		}
		c.ws.broadcast <- bytes.TrimSpace(message)
	}
}

// Write to socket
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	tableId := r.URL.RawQuery
	log.Printf("%#v", tableId)
	if ws, ok := Servers[tableId]; ok {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("couldn't upgrade connection: %#v", err)
			return
		}
		client := &Client{ws: ws, conn: conn, send: make(chan []byte, 256)}
		client.ws.register <- client

		go client.readPump()
		go client.writePump()
	} else {
		log.Println("Bad table id")
	}
}
