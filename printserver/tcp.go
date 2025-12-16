package main

import (
	"io"
	"log"
	"net"
)

func TCPServer(addr string) {
	log.Printf("TCP: Listening on %s...", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("TCP: Failed to accept connection: %v", err)
			continue
		}
		go handleTCPConn(c)
	}
}

func handleTCPConn(c net.Conn) {
	packet := readTCPConn(c)
	err := makePrint(string(packet))
	if err != nil {
		log.Printf("TCP: Failed to print: %v", err)
	}
}

func readTCPConn(c net.Conn) []byte {
	log.Printf("TCP: Connection from %s", c.RemoteAddr())
	defer c.Close()
	packet := make([]byte, 0, 4096)
	for {
		recv := make([]byte, 4096)
		n, err := c.Read(recv)
		if err != nil {
			if err != io.EOF {
				log.Printf("TCP: Failed to read connection: %v", err)
			}
			break
		}
		packet = append(packet, recv[:n]...)
	}
	return packet
}
