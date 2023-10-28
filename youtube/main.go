package main

import (
	delivery "youtube/delivery"
	repo "youtube/repository"
	uc "youtube/usecases"
)

const KEY = "AIzaSyBit2E5eTkovj4Y87AFsBkgNjXGauYjRG4"

func main() {
	repo := repo.NewYoutubeRepository(KEY)
	usecases := uc.NewYoutubeUsecases(repo)

	delivery.NewGRPCServer(usecases)
}
