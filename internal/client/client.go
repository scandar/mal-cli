package client

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/scandar/mal-cli/internal/secrets"
)

var c http.Client
var baseURL = "https://api.myanimelist.net/v2"

func initRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, baseURL+url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", "Bearer "+secrets.Get("access_token"))
	return req, nil
}

func Get(url string, params map[string]string) (*http.Response, error) {
	req, err := initRequest("GET", url, nil)
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

func Patch(reqUrl string, params map[string]string) (*http.Response, error) {
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}
	body := bytes.NewBufferString(data.Encode())

	req, err := initRequest("PATCH", reqUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return c.Do(req)
}

func Delete(url string) (*http.Response, error) {
	req, err := initRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func main() {
	c = http.Client{Timeout: time.Duration(1) * time.Second}
}
