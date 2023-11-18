package main

import (
	"images/internal/config"
	"images/internal/secrets"
	"images/pkg/delivery/grpcServer"
	"images/pkg/httpClient"
	"images/pkg/repo/unsplash"
	"images/pkg/usecases"
	"log"

	"google.golang.org/grpc"
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

	grpcServerCh := make(chan *grpc.Server)

	go func() {
		s, err := grpcServer.New(uc, cfg.GRPCPort)

		if err != nil {
			log.Printf("%v", err)
		}

		grpcServerCh <- s

	}()

	newGRPCServer := <-grpcServerCh

	defer newGRPCServer.Stop()

	select {} // block forever
}
