package repo

import (
	"encoding/json"
	"fmt"
	"apiClient"
	"net/http"
	models "images/models"
)

type resp struct {
	Urls struct{ Full string }
	Id   string
}

type myRepo struct {
	UnsplashBaseUrl  string
	UnsplashClientId string
}

func NewRepo(unsplashBaseUrl, unsplashClientId string) models.ImagesRepository {
	return &myRepo{
		UnsplashBaseUrl:  unsplashBaseUrl,
		UnsplashClientId: unsplashClientId,
	}
}

func (c *myRepo) GetImgByQuery(query string) (respUrl string, id string, err error) {
	url := fmt.Sprintf(c.UnsplashBaseUrl + "/photos/random")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("client_id", c.UnsplashClientId)
	q.Add("query", query)

	req.URL.RawQuery = q.Encode()

	api := apiClient.NewHttpClient()

	bytes, err := api.DoRequest(req)
	
	if err != nil {
		return respUrl, id, err
	}

	var data resp
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return respUrl, id, err
	}

	respUrl = data.Urls.Full
	id = data.Id

	return respUrl, id, err
}
