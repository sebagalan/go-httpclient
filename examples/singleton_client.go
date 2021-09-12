package examples

import (
	"time"

	"github.com/sebagalan/go-httpclient/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {

	return gohttp.NewClientBuilder().
		SetDialerContextTimeout(2 * time.Second).
		SetResponseHeaderTimeout(3 * time.Second).
		SetUserAgent("mycustom-useragent").
		Build()
}
