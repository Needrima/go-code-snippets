package pubsub

import (
	"fmt"
	"net"
	"slices"
	"sync"
)

type PubSub struct {
	mu       sync.Mutex
	channels map[string][]chan string
}

func New() *PubSub {
	return &PubSub{channels: map[string][]chan string{}}
}

func (p *PubSub) Subscribe(conn net.Conn, channelName string) {
	newSubscriber := make(chan string)

	p.mu.Lock()
	p.channels[channelName] = append(p.channels[channelName], newSubscriber)
	p.mu.Unlock()

	go func() {
		fmt.Println("new subscriber. ready to receive messages")
		for {
			val, ok := <-newSubscriber

			p.mu.Lock()
			if !ok { // channel is closed
				fmt.Fprintln(conn, "a subscriber left")
				// delete channel
				index := slices.Index(p.channels[channelName], newSubscriber)
				if index == -1 {
					return
				}
				_ = slices.Delete(p.channels[channelName], index, index+1)
				return
			}

			fmt.Fprintln(conn, val)
			p.mu.Unlock()
		}
	}()
}

func (p *PubSub) Publish(channelName, value string) {
	subscribers := p.channels[channelName]

	for _, sub := range subscribers {
		sub <- value
	}
}
