package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	defaultMaxIdleConnsPerHost   = 2
	defaultResponseHeaderTimeout = 200 * time.Millisecond
	dialerContextTimeout         = 500 * time.Millisecond
)

func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	dialContext := net.Dialer{
		Timeout: c.getDialerContextTimeout(),
	}

	c.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getMaxIdleConnsPerHost(),
			ResponseHeaderTimeout: c.getResponseHeaderTimeout(),
			DialContext:           dialContext.DialContext,
		},
	}

	return c.client
}

func (c *httpClient) getMaxIdleConnsPerHost() int {
	if c.maxIdleConnsPerHost > 0 {
		return c.maxIdleConnsPerHost
	}

	return defaultMaxIdleConnsPerHost
}

func (c *httpClient) getResponseHeaderTimeout() time.Duration {

	if c.disableTimeout {
		return 0
	}

	if c.responseHeaderTimeout > 0 {
		return c.responseHeaderTimeout
	}

	return defaultResponseHeaderTimeout
}

func (c *httpClient) getDialerContextTimeout() time.Duration {

	if c.disableTimeout {
		return 0
	}

	if c.dialerContextTimeout > 0 {
		return c.dialerContextTimeout
	}

	return dialerContextTimeout
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	var requestBody []byte
	var requestBodyErrors error

	requestHeaders := c.getRequestHeader(headers)

	if body != nil {
		requestBody, requestBodyErrors = c.getRequestBody(requestHeaders.Get("Content-Type"), body)
		if requestBodyErrors != nil {
			errors.New("http client fail - content type not found")
			return nil, requestBodyErrors
		}
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		errors.New("http client fail")
		return nil, err
	}

	request.Header = requestHeaders

	client := c.getHttpClient()

	return client.Do(request)
}

func (c *httpClient) getRequestHeader(requestHeaders http.Header) http.Header {

	result := make(http.Header)

	for header, value := range c.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {

	if body == nil {
		return nil, errors.New("content-type not present")
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
