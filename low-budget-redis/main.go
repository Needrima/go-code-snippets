package main

import (
	"fmt"
	"log"
	"low-budget-redis/cache"
	"low-budget-redis/connection"
	pubsub "low-budget-redis/pub-sub"
	"net"
)

const (
	SET       = "SET"
	GET       = "GET"
	PUBLISH   = "PUBLISH"
	SUBSCRIBE = "SUBSCRIBE"
)

func main() {
	li, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("could not connect to redis:", err)
	}

	cache := cache.New()
	pubsub := pubsub.New()

	for {
		// Accept incoming connections
		conn, err := li.Accept()
		if err != nil {
			log.Printf("could not accept connection: %v", err)
			continue
		}
		fmt.Println("New connection from", conn.RemoteAddr())
		go connection.HandleConn(conn, cache, pubsub)
	}
}
