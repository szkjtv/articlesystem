package main

import (
	"articlesystem/routers"
	"log"
	"time"
)

func main() {
	log.Println("http://localhost:3006")
	go routers.Router()
	for {
		time.Sleep(1 * time.Microsecond)
	}
}
