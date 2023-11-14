package internals

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Http struct {
	host string
}

func NewHttp(host string) *Http {
	return &Http{
		host: host,
	}
}

func (c *Http) Get(ctx context.Context, path string, response interface{}, headers *map[string]string) error {
	client := &http.Client{}

	request, err := http.NewRequest("GET", c.host+path, nil)
	if err != nil {
		return err
	}

	if headers != nil {
		for key, value := range *headers {
			request.Header.Set(key, value)
		}
	}

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

func (c *Http) Post(ctx context.Context, path string, payload interface{}, response interface{}, headers *map[string]string) error {
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
	if headers != nil {
		for key, value := range *headers {
			request.Header.Set(key, value)
		}
	}

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
