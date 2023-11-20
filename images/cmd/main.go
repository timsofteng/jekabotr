package main

import (
	"images/internal/config"
	"images/internal/secrets"
	"images/pkg/delivery/grpcServer"
	"images/pkg/delivery/httpServer"
	"images/pkg/httpClient"
	"images/pkg/repo/unsplash"
	"images/pkg/usecases"
	"log"
	"os"
	"os/signal"
)

func main() {
	s, err := secrets.New()

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	newHTTPClient := httpClient.New()

	newUnsplashRepo := unsplash.New(newHTTPClient, s.UnsplashClientID)

	uc := usecases.New(newUnsplashRepo)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		s, err := grpcServer.New(uc, cfg.GRPCPort)

		if err != nil {
			log.Printf("%v", err)
		}

		defer s.Stop()
	}()

	go func() {
		err := httpServer.New(uc, cfg.HTTPPort)

		if err != nil {
			log.Printf("%v", err)
		}

	}()

	<-quit
}
