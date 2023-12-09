package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.wmata.com/"
)

type HttpRequester struct {
	apiKey string
}

func New(apiKey string) *HttpRequester {
	return &HttpRequester{
		apiKey: apiKey,
	}
}

func (r *HttpRequester) SendHttpRequest(ctx context.Context, url string, data any) ([]byte, error) {
	var buf io.Reader = nil
	if data != nil {
		var ok bool
		buf, ok = data.(io.Reader)
		if !ok {
			var b []byte
			b, err := json.Marshal(data)
			if err != nil {
				return nil, fmt.Errorf("json.Marshal: %v", err)
			}
			buf = bytes.NewBuffer(b)
		}
	}

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, buf)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %v", err)
	}
	req.Header.Set("api_key", r.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid http status: %d", resp.StatusCode)
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	return responseBody, nil
}

func GenerateUrl(suburl string, params map[string]string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("url.Parse (%s): %v", baseURL, err)
	}

	queryParams := url.Values{}
	for k, v := range params {
		queryParams.Add(k, v)
	}
	encodedQuery := queryParams.Encode()

	if encodedQuery != "" {
		return base.String() + suburl + "?" + encodedQuery, nil
	}
	return base.String() + suburl, nil
}
