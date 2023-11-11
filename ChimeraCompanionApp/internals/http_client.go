package internals

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HttpClient struct {
	host string
}

func NewHttpClient(host string) *HttpClient {
	return &HttpClient{
		host: host,
	}
}

func (c *HttpClient) Post(path string, payload interface{}, response interface{}) error {
	client := &http.Client{}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", c.host+path, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil
}
