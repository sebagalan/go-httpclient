package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/sebagalan/go-httpclient/gohttp_mock"
)

func TestMain(m *testing.M) {

	fmt.Printf("Starting MockServer\n")
	gohttp_mock.MockServer.StartMockSever()

	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {

	t.Run("Test Error Fetching from url", func(t *testing.T) {

		gohttp_mock.MockServer.DeleteMocks()
		gohttp_mock.MockServer.AddMock(gohttp_mock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("Failed to get resource"),
		})

		endpoint, err := GetEndpoints()

		if endpoint != nil {
			t.Error("no valid endpoints expected")
		}

		if err == nil {
			t.Error("expeted error")
		}

		if err.Error() != "Failed to get resource" {
			t.Error("nonexpected message", err.Error())
		}
	})

	t.Run("Test Error Unmarshaler", func(t *testing.T) {

		gohttp_mock.MockServer.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})
		endpoint, err := GetEndpoints()

		if endpoint != nil {
			t.Error("no valid endpoints expected")
		}

		if err == nil {
			t.Error("expeted error")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal") {
			t.Error("nonexpected message")
		}

	})

	t.Run("Test No Error", func(t *testing.T) {

		gohttp_mock.MockServer.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})
		endpoint, err := GetEndpoints()

		if err != nil {
			t.Error("no errors expected")
		}

		if endpoint == nil {
			t.Error("expeted endpoints an got nil")
		}

		if endpoint.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid value")
		}

	})
}
