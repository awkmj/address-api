package client

import (
	"fmt"
	"io"
	"net/http"
)

type Client interface {
}

type Opt func(*HttpClient)

type HttpClient struct {
	url    string
	client *http.Client
}

func NewHttpClient(opts ...Opt) *HttpClient {

	httpClient := &HttpClient{
		client: &http.Client{},
	}

	for _, opt := range opts {
		opt(httpClient)
	}
	return httpClient
}

func WithUrl(url string) func(*HttpClient) {
	return func(h *HttpClient) {
		h.url = url
	}
}

func (h *HttpClient) Get(zipCode string) (io.Reader, error) {

	url := fmt.Sprintf("%s/ws/%s/json/", h.url, zipCode)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := h.client.Do(req)

	if err != nil {
		return nil, err
	}

	return res.Body, nil
}
