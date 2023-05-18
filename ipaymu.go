package ipaymu_go_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func (c Client) CallApi(url *url.URL, signature string, body []byte) {
	reqBody := ioutil.NopCloser(strings.NewReader(string(body)))

	req := &http.Request{
		Method: http.MethodPost,
		URL:    url,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"va":           {c.VirtualAccount},
			"signature":    {signature},
		},
		Body: reqBody,
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c Client) DirectPaymentVA(request RequestDirectVA) (Response, error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/direct", c.EnvApi))
	jsonBody, _ := json.Marshal(request)
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", c))
	c.CallApi(url, signature, jsonBody)

	return Response{}, nil
}

func (c Client) DirectPaymentConStore(request RequestDirectConStore) (Response, error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/direct", c.EnvApi))
	jsonBody, _ := json.Marshal(request)
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", c))
	c.CallApi(url, signature, jsonBody)

	return Response{}, nil
}

func (c Client) RedirectPayment(request RequestRedirect) (Response, error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/", c.EnvApi))
	jsonBody, _ := json.Marshal(request)
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", c))
	c.CallApi(url, signature, jsonBody)

	return Response{}, nil
}
