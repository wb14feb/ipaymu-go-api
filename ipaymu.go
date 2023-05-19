package ipaymu_go_api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	ApiKey         string
	VirtualAccount string
	EnvApi         EnvironmentType
}

func NewClient() *Client {
	return &Client{
		EnvApi: Production,
	}
}

var defHTTPTimeout = 30 * time.Second

func (c Client) CallApi(url *url.URL, signature string, body []byte) ([]byte, error) {
	reqBody := ioutil.NopCloser(strings.NewReader(string(body)))

	req := &http.Request{
		Method: http.MethodPost,
		URL:    url,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"va":           {c.VirtualAccount},
			"signature":    {signature},
			"Accept":       {"application/json"},
		},
		Body: reqBody,
	}

	httpClient := http.DefaultClient
	httpClient.Timeout = defHTTPTimeout
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Printf("An Error Occured %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return respBody, nil
}
