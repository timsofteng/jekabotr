package repository

import (
	"encoding/json"
	"fmt"
	"jekabot/apiClient"
	"jekabot/models"
	"net/http"
)

type resp struct {
	Urls struct{ Full string }
	Id   string
}

type myTaksaRepo struct {
	UnsplashBaseUrl  string
	UnsplashClientId string
}

func NewTaksaRepository(unsplashBaseUrl, unsplashClientId string) models.TaksaRepository {
	return &myTaksaRepo{
		UnsplashBaseUrl:  unsplashBaseUrl,
		UnsplashClientId: unsplashClientId,
	}
}

func (c *myTaksaRepo) GetRandomTaksaUrl() (respUrl string, id string, err error) {
	url := fmt.Sprintf(c.UnsplashBaseUrl + "/photos/random")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("client_id", c.UnsplashClientId)
	q.Add("query", "dachshund")

	req.URL.RawQuery = q.Encode()

	api := apiClient.NewHttpClient()

	bytes, err := api.DoRequest(req)
	if err != nil {
		return
	}

	var data resp
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return
	}

	respUrl = data.Urls.Full
	id = data.Id

	return
}
