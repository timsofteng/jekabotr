package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jekabot/models"
	"net/http"
	"time"
)

type Client struct {
	UnsplashBaseUrl  string
	UnsplashClientId string
}

func NewClient(unsplashBaseUrl, unsplashClientId string) models.ApiMethods {
	return &Client{
		UnsplashBaseUrl:  unsplashBaseUrl,
		UnsplashClientId: unsplashClientId,
	}
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func doRequest(req *http.Request) (body []byte, err error) {
	client := httpClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return
}

func (c *Client) GetRandomTaksaUrl() (respUrl string, id string, err error) {
	url := fmt.Sprintf(c.UnsplashBaseUrl + "/photos/random")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("client_id", c.UnsplashClientId)
	q.Add("query", "dachshund")

	req.URL.RawQuery = q.Encode()

	bytes, err := doRequest(req)
	if err != nil {
		return
	}

	var data models.RandomSingleImg
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return
	}

	respUrl = data.Urls.Full
	id = data.Id

	return
}

func (c *Client) GetBytesFromUrl(url string) (bytes []byte, err error) {
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
