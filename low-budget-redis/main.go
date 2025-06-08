package main

import (
	"fmt"
	"log"
	"low-budget-redis/cache"
	"low-budget-redis/connection"
	"low-budget-redis/database"
	pubsub "low-budget-redis/pub-sub"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("could not connect to LBR:", err)
	}

	cache := cache.New()
	pubsub := pubsub.New()

	database := database.InitializeFileStorage()
	database.LoadUpDataHistoryIntoCache(cache)

	fmt.Println("LBR fired up—running lean and mean!🚀🔥💪")

	for {
		// Accept incoming connections
		conn, err := li.Accept()
		if err != nil {
			log.Printf("could not accept connection: %v", err)
			continue
		}
		fmt.Println("New connection from", conn.RemoteAddr())
		go connection.HandleConn(conn, cache, pubsub, database)
	}
}
