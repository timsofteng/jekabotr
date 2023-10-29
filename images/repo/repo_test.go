package repo

import (
	secrets "images/secrets"
	"testing"
)

// TestGetImgByQuery calls repo.TestGetImgByQuery with a some query, checking
// for a valid return value.
func TestGetImgByQuery(t *testing.T) {
	repo := NewRepo(secrets.BaseURL, secrets.ClientID)

	url, id, err := repo.GetImgByQuery("ball")

	if err != nil {
		t.Errorf(`Failed to get image %v`, err)
	}

	if len(id) < 1 {
		t.Errorf(`id is empty`)
	}

	if len(url) < 1 {
		t.Errorf(`url is empty`)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// func TestHelloEmpty(t *testing.T) {
// 	msg, err := Hello("")
// 	if msg != "" || err == nil {
// 		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
// 	}
// }
