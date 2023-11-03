package repo

import (
	secrets "images/secrets"
	"testing"
)

// TestGetImgByQuery calls repo.TestGetImgByQuery with a some query, checking
// for a valid return value.
func TestGetImgByQuerySuccess(t *testing.T) {
	repo := NewRepo(secrets.BaseURL, secrets.ClientID)

	url, id, _ := repo.ImgByQueryFetcher("ball")

	if len(id) < 1 {
		t.Errorf(`id is empty`)
	}

	if len(url) < 1 {
		t.Errorf(`url is empty`)
	}
}

func TestGetImgByQueryError(t *testing.T) {
	repo := NewRepo("", secrets.ClientID)

	_, _, err := repo.ImgByQueryFetcher("ball")

	if err == nil {
		t.Errorf(`Bad query %v`, err)
	}
}
