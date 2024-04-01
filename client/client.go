package client

import (
	"net/http"
	"time"

	"github.com/scandar/mal-cli/secrets"
)

var c http.Client
var baseURL = "https://api.myanimelist.net/v2"
var URLs = map[string]string{
	"userInfo": "/users/@me",
}

func initRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, baseURL+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", "Bearer "+secrets.Get("access_token"))
	return req, nil
}

func Get(url string, params map[string]string) (*http.Response, error) {
	req, err := initRequest("GET", url)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	return c.Do(req)
}

func main() {
	c = http.Client{Timeout: time.Duration(1) * time.Second}
}
