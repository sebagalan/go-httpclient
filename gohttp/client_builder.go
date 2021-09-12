package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	maxIdleConnsPerHost   int
	responseHeaderTimeout time.Duration
	dialerContextTimeout  time.Duration
	headers               http.Header
	disableTimeout        bool
	userAgent             string

	client *http.Client
}

type ClientBuilder interface {
	SetMaxIdleConnsPerHost(int) ClientBuilder
	SetResponseHeaderTimeout(time.Duration) ClientBuilder
	SetDialerContextTimeout(time.Duration) ClientBuilder
	DisableTimeout(bool) ClientBuilder
	SetHeaders(http.Header) ClientBuilder
	SetHttpClient(client *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder

	Build() Client
}

func (c *clientBuilder) Build() Client {

	client := &httpClient{
		builder: c,
	}

	return client
}

func NewClientBuilder() ClientBuilder {

	clientBuilder := &clientBuilder{
		headers:        make(http.Header),
		disableTimeout: false,
	}

	clientBuilder.headers.Set("Accept", "application/json")
	clientBuilder.headers.Set("Content-Type", "application/json")

	return clientBuilder
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetMaxIdleConnsPerHost(maxIdleConnsPerHost int) ClientBuilder {
	c.maxIdleConnsPerHost = maxIdleConnsPerHost
	return c
}

func (c *clientBuilder) SetResponseHeaderTimeout(responseHeaderTimeout time.Duration) ClientBuilder {
	c.responseHeaderTimeout = responseHeaderTimeout
	return c
}

func (c *clientBuilder) SetDialerContextTimeout(dialerContextTimeout time.Duration) ClientBuilder {
	c.dialerContextTimeout = dialerContextTimeout
	return c
}

func (c *clientBuilder) DisableTimeout(disable bool) ClientBuilder {
	c.disableTimeout = disable
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	return c
}
