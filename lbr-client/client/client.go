package client

import (
	"bufio"
	"fmt"
	"net"
)

const (
	SET       = "SET"
	GET       = "GET"
	PUBLISH   = "PUBLISH"
	SUBSCRIBE = "SUBSCRIBE"
)

type Client struct {
	conn net.Conn
}

func New(port string) (*Client, error) {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("could not dial in to port %v: %v",port, err)
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Set(key, value string) {
	command := fmt.Sprintf("%s %s %s", SET, key, value)
	fmt.Fprintln(c.conn, command)
}

func (c *Client) Get(key string) string {
	responseChan := make(chan string, 1)
	go func() {
		scanner := bufio.NewScanner(c.conn)
		scanner.Scan()
		text := scanner.Text()
		responseChan <- text
	}()

	command := fmt.Sprintf("%s %s",GET, key)
	fmt.Fprintln(c.conn, command)

	return <-responseChan
}

func (c *Client) Publish(channelName, message string) {
	command := fmt.Sprintf("%s %s %s", PUBLISH, channelName, message)
	fmt.Fprintln(c.conn, command)
}

func (c *Client) Subscribe(channelName string) <-chan string {
	responseChan := make(chan string, 1)
	go func() {
		scanner := bufio.NewScanner(c.conn)
		for scanner.Scan() {
			text := scanner.Text()
			responseChan <- text
		}
	}()

	command := fmt.Sprintf("%s %s", SUBSCRIBE, channelName)
	fmt.Fprintln(c.conn, command)

	return responseChan
}
