package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pubChannel := "channel1"
	bs, err := json.Marshal(User{FirstName: "Jon", LastName: "Snow"})
	if err  != nil {
		log.Println("marshalling user:", err)
		return
	}
	data := string(bs)

	if err := client.Publish(context.Background(), pubChannel, data).Err(); err != nil {
		log.Fatal("error publishing to channel:", err)
	}

	log.Println("published message to channel")
}
