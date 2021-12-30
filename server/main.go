package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8081", "http service address")

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	serv := gin.Default()
	serv.SetTrustedProxies([]string{"192.168.0.0/24"})

	serv.GET("/new_game", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "localhost:3000")
		tableId := generateTableId(3)
	})
}
