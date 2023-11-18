package secrets

import (
	"errors"
	"os"
)

type secrets struct {
	UnsplashClientID string
}

func New() (*secrets, error) {
	unsplashClientID := os.Getenv("UNSPLASH_CLIENT_ID")

	if len(unsplashClientID) < 1 {
		return nil, errors.New("no unsplash client id")
	}

	return &secrets{UnsplashClientID: unsplashClientID}, nil

}
