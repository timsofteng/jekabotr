package repo

import (
	"apiClient"
	"encoding/json"
	"fmt"
	"net/http"
)

type resp struct {
	Urls struct{ Full string }
	Id   string
}

type myRepo struct {
	UnsplashBaseUrl  string
	UnsplashClientId string
}

func NewRepo(unsplashBaseUrl, unsplashClientId string) *myRepo {
	return &myRepo{
		UnsplashBaseUrl:  unsplashBaseUrl,
		UnsplashClientId: unsplashClientId,
	}
}

func (c myRepo) ImgByQueryFetcher(query string) (bin []byte, id string, err error) {
	url := fmt.Sprintf(c.UnsplashBaseUrl + "/photos/random")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, id, err
	}

	q := req.URL.Query()
	q.Add("client_id", c.UnsplashClientId)
	q.Add("query", query)

	req.URL.RawQuery = q.Encode()

	api := apiClient.NewHttpClient()

	bin, err = api.DoRequest(req)

	if err != nil {
		return nil, id, err
	}

	var data resp
	err = json.Unmarshal(bin, &data)

	if err != nil {
		return nil, id, err
	}

	respUrl := data.Urls.Full
	id = data.Id

	req, err = http.NewRequest("GET", respUrl, nil)

	if err != nil {
		return nil, id, err
	}

	q = req.URL.Query()
	q.Add("client_id", c.UnsplashClientId)
	req.URL.RawQuery = q.Encode()

	bin, err = api.DoRequest(req)

	if err != nil {
		return nil, id, err
	}

	return bin, id, err
}
