package main

import (
	delivery "images/delivery"
	repo "images/repo"
	secrets "images/secrets"
	uc "images/usecases"
)

func main() {
	taksaRepo := repo.NewRepo(secrets.BaseURL, secrets.ClientID)

	taksaUs := uc.NewTaksaUsecases(taksaRepo)

	delivery.NewGRPCServer(taksaUs)
}
