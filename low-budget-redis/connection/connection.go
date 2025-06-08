package connection

import (
	"bufio"
	"fmt"
	"low-budget-redis/cache"
	"low-budget-redis/database"
	pubsub "low-budget-redis/pub-sub"
	"net"
	"strings"
	"sync"
)

const (
	SET       = "SET"
	GET       = "GET"
	PUBLISH   = "PUBLISH"
	SUBSCRIBE = "SUBSCRIBE"
)

func HandleConn(conn net.Conn, cache *cache.Cache, pubsub *pubsub.PubSub, database *database.Database) {
	var mu sync.Mutex

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.ToLower(text) == "exit" {
			fmt.Fprintln(conn, "closing connection...")
			conn.Close()
			return
		}

		fields := strings.Fields(text)
		if len(fields) == 0 {
			fmt.Fprintln(conn, "empty command")
			continue
		}
		cmd := fields[0]

		switch cmd {
		case SET:
			if len(fields) < 3 {
				fmt.Fprintf(conn, "invalid number of arguments %v\n", text)
				continue
			}
			key, value := fields[1], strings.Join(fields[2:], " ")

			mu.Lock()
			command := fmt.Sprintf("%s %s %s", cmd, key, value)
			database.Insert(command)
			mu.Unlock()

			cache.Set(key, value)

		case GET:
			if len(fields) != 2 {
				fmt.Fprintf(conn, "invalid number of arguments %v\n", text)
				continue
			}
			value, ok := cache.Get(fields[1])
			if !ok {
				fmt.Fprintf(conn, "key %v does not exist\n", fields[1])
				continue
			}

			fmt.Fprintln(conn, value)

		case SUBSCRIBE:
			if len(fields) != 2 {
				fmt.Fprintf(conn, "invalid number of arguments %v\n", text)
				continue
			}
			channelName := fields[1]
			pubsub.Subscribe(conn, channelName)

		case PUBLISH:
			if len(fields) < 3 {
				fmt.Fprintf(conn, "invalid number of arguments %v\n", text)
				continue
			}
			channelName, value := fields[1], strings.Join(fields[2:], " ")
			pubsub.Publish(channelName, value)

		default:
			fmt.Fprintf(conn, "invalid command %v\n", cmd)
		}
	}
}
