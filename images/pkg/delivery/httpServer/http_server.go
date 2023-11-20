package httpServer

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type randImgResp struct {
	Url string `json:"url"`
}

type usecases interface {
	RandomImg(ctx context.Context, query string) (url string, err error)
}

func randImgHandler(uc usecases) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		log.Printf("%v", query)
		if query == "" {
			err := errors.New("no query")
			log.Printf("%v", err)
			return
		}

		ctx := context.Background()

		url, err := uc.RandomImg(ctx, query)

		resp := randImgResp{Url: url}

		w.Header().Set("Content-Type", "application/json")

		jsonResp, err := json.Marshal(resp)

		_, err = w.Write(jsonResp)

		if err != nil {
			log.Printf("%v", err)
			return
		}

	})

}

func New(uc usecases, portNum string) error {
	log.Println("Started http server on port", portNum)

	// Spinning up the server.

	randImgHandler(uc)
	err := http.ListenAndServe(portNum, nil)

	return err
}
