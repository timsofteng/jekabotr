package main

import (
	"log"
	"os"
	"youtube/delivery"
	"youtube/repo"
	uc "youtube/usecases"
)

func main() {
	apiKey := os.Getenv("API_KEY")

	if len(apiKey) < 1 {
		log.Fatal("No api key")
	}

	repo := repo.NewYoutubeRepository(apiKey)
	usecases := uc.NewYoutubeUsecases(repo)

	err := delivery.NewGRPCServer(usecases)
	if err != nil {
		log.Printf("%v", err)
	}
}
