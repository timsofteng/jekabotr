package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	config, err := ReadConfig()

	if err != nil {
		log.Fatalf("error to read config, %e", err)
	}

	ytClient, ytGRPCConn, err := NewYoutubeGRPCClient()

	defer ytGRPCConn.Close()

	if err != nil {
		log.Fatalf("failed to start yt grpc client, %e", err)
	}

	d := NewDelivery(
		ytClient,
		config,
	)

	for update := range d.Updates {
		if update.Message != nil { // If we got a message
			go d.Router(update)
		}
	}

}
