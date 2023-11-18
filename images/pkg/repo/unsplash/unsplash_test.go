package unsplash

import (
	"context"
	"images/pkg/httpClient"
	secrets "images/internal/secrets"
	"testing"
	"time"
)

// TestGetImgByQuery calls repo.TestGetImgByQuery with a some query, checking
// for a valid return value.
func TestGetImgByQuerySuccess(t *testing.T) {
	s, err := secrets.New()

	if err != nil {
		t.Errorf(`secrets are empty`)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))
	defer cancel()

	newHTTPClient := httpClient.New()

	repo := New(newHTTPClient, s.UnsplashClientID)

	url, _ := repo.ImgByQueryFetcher(ctx, "ball")

	if len(url) < 1 {
		t.Errorf(`url is empty`)
	}
}

// func TestGetImgByQueryError(t *testing.T) {
// 	repo := NewRepo("", secrets.ClientID)

// 	_, _, err := repo.ImgByQueryFetcher("ball")

// 	if err == nil {
// 		t.Errorf(`Bad query %v`, err)
// 	}
// }
