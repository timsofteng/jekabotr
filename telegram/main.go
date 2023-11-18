package main

import (
	"log"
	"telegram/delivery"
	"telegram/repo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	config, err := ReadConfig()

	if err != nil {
		log.Fatalf("error to read config, %e", err)
	}

	ytClient, ytGRPCConn, err := repo.NewYoutubeGRPCClient()
	if err != nil {
		log.Fatalf("failed to start yt grpc client, %e", err)
	}

	imagesClient, imagesGRPCConn, err := repo.NewImagesGRPCClient()
	if err != nil {
		log.Fatalf("failed to start images grpc client, %e", err)
	}

	defer ytGRPCConn.Close()
	defer imagesGRPCConn.Close()

	d, err := delivery.NewDelivery(
		ytClient,
		imagesClient,
		config,
	)

	if err != nil {
		log.Printf("%v", err)
	}

	for update := range d.Updates {
		if update.Message != nil { // If we got a message
			go func(update tgbotapi.Update) {
				err := d.Router(update)
				if err != nil {
					log.Printf("%v", err)
				}
			}(update)
		}
	}

}
