package utils

import (
	"io"
	"net/http"
)

func BytesFromUrl(url string) (bytes []byte, err error) {
	response, err := http.Get(url)

	if err != nil {
		return
	}

	bytes, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		// log.Print("img request not 200")
		return
	}

	return
}
