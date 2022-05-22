package singleton

import (
	"net/http"
	"sync"
)

type HttpClient struct {
	Id string
	http.Client
}

var httpClient *HttpClient
var httpInitOnce sync.Once

func NewHttpClient() *HttpClient {
	if httpClient == nil {
		httpInitOnce.Do(func() {
			httpClient = &HttpClient{
				Id: "1",
			}
		})
	}
	return httpClient
}
