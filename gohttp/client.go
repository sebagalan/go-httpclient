package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client                *http.Client
	maxIdleConnsPerHost   int
	responseHeaderTimeout time.Duration
	dialerContextTimeout  time.Duration
	headers               http.Header
	disableTimeout        bool
}

func NewHttpClient() HttpClient {

	httpClient := &httpClient{
		headers: make(http.Header),
	}

	httpClient.headers.Set("Accept", "application/json")
	httpClient.headers.Set("Content-Type", "application/json")

	return httpClient
}

type HttpClient interface {
	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header) (*http.Response, error)

	SetMaxIdleConnsPerHost(int)
	SetResponseHeaderTimeout(time.Duration)
	SetDialerContextTimeout(time.Duration)
	DisableTimeout(bool)
}

func SetHeaders() {}

func (c *httpClient) SetMaxIdleConnsPerHost(maxIdleConnsPerHost int) {
	c.maxIdleConnsPerHost = maxIdleConnsPerHost
}

func (c *httpClient) SetResponseHeaderTimeout(responseHeaderTimeout time.Duration) {
	c.responseHeaderTimeout = responseHeaderTimeout
}

func (c *httpClient) SetDialerContextTimeout(dialerContextTimeout time.Duration) {
	c.dialerContextTimeout = dialerContextTimeout
}

func (c *httpClient) DisableTimeout(disable bool) {
	c.disableTimeout = disable
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
