package main

import (
	"log"
	repo "telegram/repo"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	config, err := ReadConfig()

	if err != nil {
		log.Fatalf("error to read config, %e", err)
	}

	ytClient, ytGRPCConn, err := repo.NewYoutubeGRPCClient()
	imagesClient, imagesGRPCConn, err := repo.NewImagesGRPCClient()

	defer ytGRPCConn.Close()
	defer imagesGRPCConn.Close()

	if err != nil {
		log.Fatalf("failed to start yt grpc client, %e", err)
	}

	d := NewDelivery(
		ytClient,
		imagesClient,
		config,
	)

	for update := range d.Updates {
		if update.Message != nil { // If we got a message
			go d.Router(update)
		}
	}

}
