package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	subChannel := "channel1"

	sub1 := client.Subscribe(context.Background(), subChannel)

	subCh := sub1.Channel()
	data := map[string]interface{}{}

	go func(ch <-chan *redis.Message) {
		for {
			msg := <-ch
			fmt.Println("recieved message:", msg.Payload)
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				log.Println("unmarshal error:", err)
			}

			fmt.Printf("Firstname: %v\nLastname: %v\n", data["first_name"], data["last_name"])
		}
	}(subCh)
	
	fmt.Println("subscribing to channel")
	select {}
}
