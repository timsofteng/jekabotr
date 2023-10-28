package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	c, err := ReadConfig()

	if err !=nil {
		log.Fatalf("error to read config, %e", err)
	}

	d := NewDelivery(
		c,
	)

	for update := range d.Updates {
		if update.Message != nil { // If we got a message
			go d.Router(update)
		}
	}

}
