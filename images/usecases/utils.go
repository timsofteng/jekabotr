package usecases

import (
	"io"
	"net/http"
)

func BytesFromUrl(url string) (bin []byte, err error) {
	response, err := http.Get(url)

	if err != nil {
		return bin, err
	}

	bin, err = io.ReadAll(response.Body)
	if err != nil {
		return bin, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		// log.Print("img request not 200")
		return bin, err
	}

	return bin, err
}
