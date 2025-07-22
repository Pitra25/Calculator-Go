package connection

import (
	"net/http"
)

type ApiCLient struct {
	BaseURL string
	Client  *http.Client
}

func NewApiCLient(baseUrl string) *ApiCLient {
	return &ApiCLient{
		BaseURL: baseUrl,
		Client:  &http.Client{},
	}
}
