package unsplash

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

var baseURL = "https://api.unsplash.com"

type resp struct {
	Urls struct {
		Full string `json:"full" validate:"required"`
	}
	Id string `json:"id" validate:"required"`
}

type HTTPClient interface {
	DoRequest(req *http.Request) (reps []byte, err error)
}

type myRepo struct {
	validate         *validator.Validate
	httpClient       HTTPClient
	unsplashClientID string
}

func New(httpClient HTTPClient, unsplashClientID string) *myRepo {
	validate := validator.New()

	return &myRepo{
		httpClient:       httpClient,
		validate:         validate,
		unsplashClientID: unsplashClientID,
	}
}

func (c myRepo) ImgByQueryFetcher(ctx context.Context, query string) (url string, err error) {

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/photos/random", nil)

	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("client_id", c.unsplashClientID)
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	jsonBin, err := c.httpClient.DoRequest(req)

	if err != nil {
		return "", err
	}

	var data resp
	err = json.Unmarshal(jsonBin, &data)

	if err != nil {
		return "", err
	}

	err = c.validate.Struct(data)

	if err != nil {
		return "", err
	}

	url = data.Urls.Full

	return url, err

}

// func (c myRepo) getBinFromURL(ctx context.Context, fileURL string) ([]byte, error) {
// 	req := c.request
// 	req.Method = "GET"
// 	req.URL.Path = fileURL

// 	bin, err := c.httpClient.DoRequest(req)

// 	if err != nil {
// 		return nil, err
// 	}

// 	err = c.validate.Struct(bin)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return bin, err
// }
