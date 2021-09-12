package gohttp

import "net/http"

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
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

	if c.builder.userAgent != nil {
		if result.Get("User-Agent") != "" {
			return result
		}
		result.Set("User-Agent", c.builder.userAgent)
	}

	return result
}
