package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	defaultMaxIdleConnsPerHost   = 5
	defaultResponseHeaderTimeout = 2001 * time.Millisecond
	dialerContextTimeout         = 5 * time.Millisecond
)

func (c *httpClient) getHttpClient() *http.Client {

	c.doOnce.Do(func() {
		dialContext := net.Dialer{
			Timeout: c.getDialerContextTimeout(),
		}

		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}

		c.client = &http.Client{
			Timeout: c.getDialerContextTimeout() + c.getResponseHeaderTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnsPerHost(),
				ResponseHeaderTimeout: c.getResponseHeaderTimeout(),
				DialContext:           dialContext.DialContext,
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnsPerHost() int {
	if c.builder.maxIdleConnsPerHost > 0 {
		return c.builder.maxIdleConnsPerHost
	}

	return defaultMaxIdleConnsPerHost
}

func (c *httpClient) getResponseHeaderTimeout() time.Duration {

	if c.builder.disableTimeout {
		return 0
	}

	if c.builder.responseHeaderTimeout > 0 {
		return c.builder.responseHeaderTimeout
	}

	return defaultResponseHeaderTimeout
}

func (c *httpClient) getDialerContextTimeout() time.Duration {

	if c.builder.disableTimeout {
		return 0
	}

	if c.builder.dialerContextTimeout > 0 {
		return c.builder.dialerContextTimeout
	}

	return dialerContextTimeout
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {
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

	if mock := mockServer.getMock(method, url, string(requestBody)); mock != nil {
		return mock.GetResponse()
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		errors.New("http client fail")
		return nil, err
	}

	request.Header = requestHeaders

	client := c.getHttpClient()

	stdHttpResponse, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer stdHttpResponse.Body.Close()
	stdBodyBytes, _ := ioutil.ReadAll(stdHttpResponse.Body)

	response := &Response{
		status:     stdHttpResponse.Status,
		statusCode: stdHttpResponse.StatusCode,
		headers:    stdHttpResponse.Header,
		body:       stdBodyBytes,
	}

	return response, nil
}

func (c *httpClient) getRequestHeader(requestHeaders http.Header) http.Header {

	result := make(http.Header)

	for header, value := range c.builder.headers {
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
