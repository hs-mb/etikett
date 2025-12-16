package main

import (
	"context"
	"log"
	"net/http"

	"github.com/coder/websocket"
)

func WebSocketServer(addr string) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: false,
			OriginPatterns: []string{"*"},
		})
		if err != nil {
			log.Printf("WS: Failed to accept connection: %v", err)
			return
		}

		log.Printf("WS: Connection from %s", r.RemoteAddr)
		go handleWSConn(c)
	})
	log.Printf("WS: Listening on %s...", addr)
	http.ListenAndServe(addr, handler)
}

func handleWSConn(c *websocket.Conn) {
	defer c.CloseNow()
	packet := readWSConn(c)
	err := makePrint(string(packet))
	if err != nil {
		log.Printf("WS: Failed to print: %v", err)
	}
}

func readWSConn(c *websocket.Conn) []byte {
	packet := make([]byte, 0)
	ctx := context.Background()
	for {
		_, recv, err := c.Read(ctx)
		if err != nil {
			break
		}
		packet = append(packet, recv...)
	}
	return packet
}
