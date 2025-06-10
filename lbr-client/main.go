package main

import (
	// "fmt"
	"lbr-client/client"
	"log"
)

func main() {
	client, err := client.New(":3000")
	if err != nil {
		log.Fatal(err)
	}

	client.Set("foo", "bar")

	// value:= client.Get("foo")
	// fmt.Println(value)

	// messagesChan := client.Subscribe("foo")
	// for val := range messagesChan{
	// 	fmt.Println(val)
	// }

	// client.Publish("foo", "it's go time") // open another connection with telnet in your terminal
}
