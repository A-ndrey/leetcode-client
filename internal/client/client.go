package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const url = "https://leetcode.com/graphql"

var httpClient http.Client

func init() {
	httpClient = http.Client{
		Timeout: 30 * time.Second,
	}
}

func Do(request Request, respData any) (*Response, error) {
	req, err := makeHTTPReq(request)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return makeLeetCodeResp(resp, respData)
}

func makeHTTPReq(request Request) (*http.Request, error) {
	m, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(m)

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func makeLeetCodeResp(response *http.Response, respData any) (*Response, error) {
	resp := Response{Data: respData}
	err := json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
