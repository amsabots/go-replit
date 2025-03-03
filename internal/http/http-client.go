package http

import (
	"net/http"
	"time"
)


type HTTPClient struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) HTTPClient{
	return HTTPClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}