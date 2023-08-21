package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jekabot/models"
	"jekabot/apiClient"
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

func (c *myTaksaRepo) GetBytesFromUrl(url string) (bytes []byte, err error) {
	response, err := http.Get(url)

	if err != nil {
		return
	}

	bytes, err = ioutil.ReadAll(response.Body)
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
