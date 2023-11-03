package apiClient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type myHttpClient struct {
	ApiClient *http.Client
}

func NewHttpClient() *myHttpClient {
	client := &http.Client{Timeout: 10 * time.Second}
	return &myHttpClient{ApiClient: client}
}


func (a *myHttpClient) DoRequest(req *http.Request) (body []byte, err error) {
	resp, err := a.ApiClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
