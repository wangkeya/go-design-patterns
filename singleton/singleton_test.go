package singleton

import (
	"net/http"
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	for i := 0; i < 100; i++ {
		client := NewHttpClient()
		res, err := client.Get("https://www.baidu.com")
		if err != nil {
			t.Error("http get failed")
		}
		if res.StatusCode != http.StatusOK {
			t.Error("http get status code failed")
		}
	}

}
