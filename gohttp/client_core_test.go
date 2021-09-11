package gohttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestHeader(t *testing.T) {

	headers := make(http.Header)

	headers.Set("X-test-header", "value")

	client := httpClient{
		builder: &clientBuilder{},
	}

	finalHeaders := client.getRequestHeader(headers)

	if len(finalHeaders) == 0 {
		t.Error("default headers are not setings")
	}

	if finalHeaders.Get("X-test-header") != "value" {
		t.Error("header typo can be occured")
	}

	if finalHeaders.Get("X-testHeader") != "" {
		t.Error("header typo can be occured")
	}

}

func TestGetRequestBody(t *testing.T) {

	headers := make(http.Header)

	client := httpClient{
		builder: &clientBuilder{},
	}

	t.Run("Nil body", func(t *testing.T) {
		body, err := client.getRequestBody(headers.Get("Content-type"), nil)

		if err == nil {
			t.Error("expected no error")
		}

		if body != nil {
			t.Error("expected nil")
		}
	})

	t.Run("json body", func(t *testing.T) {
		headers.Set("Content-type", "application/json")
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody(headers.Get("Content-type"), requestBody)

		if err != nil {
			t.Error("expected no error")
		}

		if body == nil {
			t.Error("expected not nil")
		}

		if string(body) != `["one","two"]` {
			t.Error("expected json")
			fmt.Println(string(body))
		}

	})

	t.Run("xml body", func(t *testing.T) {
		headers.Set("Content-type", "application/xml")
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody(headers.Get("Content-type"), requestBody)

		if err != nil {
			t.Error("expected no error")
		}

		if body == nil {
			t.Error("expected not nil")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("expected json")
			fmt.Println(string(body))
		}
	})

}
