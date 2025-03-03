package httpclient

import (
	"net/http"
	"time"
)


type HTTPClient struct {
	HTTP *http.Client
}

func NewClient(timeout time.Duration) HTTPClient{
	return HTTPClient{
		HTTP: &http.Client{
			Timeout: timeout,
		},
	}
}