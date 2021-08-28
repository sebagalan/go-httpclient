package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sebagalan/go-httpclient/gohttp"
)

var (
	gitHttpClient   = gohttp.NewHttpClient()
	localHttpClient = gohttp.NewHttpClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getLocalGreeting() {
	headers := make(http.Header)
	headers.Add("Authorization", "Bearer ABC-123")

	response, err := localHttpClient.Get("http://localhost:8080", headers)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

func getUrl() {

	response, err := gitHttpClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

func postUrl(user User) {

	response, err := localHttpClient.Post("http://localhost:8080", nil, user)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))

}

func main() {

	getUrl()
	getLocalGreeting()

	user := User{
		FirstName: "juan",
		LastName:  "Galan",
	}

	postUrl(user)
}
