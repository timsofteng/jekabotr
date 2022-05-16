package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jekabot/models"
	"log"
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

func doRequest(req *http.Request) ([]byte, error) {
	client := httpClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (c *Client) GetRandomTaksaUrl() (string, string, error) {
	url := fmt.Sprintf(c.UnsplashBaseUrl + "/photos/random")

	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("client_id", c.UnsplashClientId)
	q.Add("query", "dachshund")

	req.URL.RawQuery = q.Encode()

	if err != nil {
		log.Print("request error")
	}

	bytes, err := doRequest(req)
	if err != nil {
	}

	var data models.RandomSingleImg
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Print("json parse error")
	}

	respUrl := data.Urls.Full
	id := data.Id

	return respUrl, id, err
}

func (c *Client) GetBytesFromUrl(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		log.Print("img request error")
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print("error due converting resp to byte")
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Print("img request not 200")
	}

	return bodyBytes, err
}
