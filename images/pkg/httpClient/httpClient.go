package httpClient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type myHTTPClient struct {
	client *http.Client
}

func New() myHTTPClient {
	return myHTTPClient{
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (a myHTTPClient) DoRequest(req *http.Request) (body []byte, err error) {
	resp, err := a.client.Do(req)
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
