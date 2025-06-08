package database

import (
	"bufio"
	"log"
	"low-budget-redis/cache"
	"os"
	"strings"
)

type Database struct {
	file *os.File
}

func InitializeFileStorage() *Database {
	f, err := os.OpenFile("db.store", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("could not open database:", err)
	}

	return &Database{file: f}
}

func (d *Database) LoadUpDataHistoryIntoCache(cache *cache.Cache) {
	scanner := bufio.NewScanner(d.file)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)

		key, value := fields[1], strings.Join(fields[2:], " ")
		cache.Set(key, value)
	}
}

func (d *Database) Insert(command string) {
	d.file.WriteString(command + "\n")
}
