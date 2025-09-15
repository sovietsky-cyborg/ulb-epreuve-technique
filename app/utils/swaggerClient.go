package utils

import (
	"io"
	"net/http"
	_ "os"
)

type SwaggerClient struct {
	url string
}

var client *SwaggerClient

func init() {
	client = &SwaggerClient{
		url: GetEnv("SWAGGER_CLIENT_URL"),
	}
}

func GetClient() *SwaggerClient {
	return client
}

func (c *SwaggerClient) GetData(path string) ([]byte, error) {

	resp, err := http.Get(c.url + path)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}
