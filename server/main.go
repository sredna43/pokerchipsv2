package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8081", "http service address")

var (
	Servers = make(map[string]*Ws)
)

type NewTableResponse struct {
	Id string `json:"id"`
}

type WsErrorResponse struct {
	Message string `json:"message"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func generateTableId(n int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rng.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	flag.Parse()
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.0.0/24"})

	Servers["test"] = NewServer("test")
	go Servers["test"].run()

	router.GET("/new_game", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "localhost:3000")
		tableId := generateTableId(3)
		Servers[tableId] = NewServer(tableId)
		if ws, ok := Servers[tableId]; ok {
			go ws.run()
		}
		c.JSON(200, NewTableResponse{Id: tableId})
	})

	router.GET("/ws", func(c *gin.Context) {
		serveWs(c.Writer, c.Request)
	})

	router.Run(*addr)
}
