package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {

	response, err := httpClient.Get("https://api.github.com")

	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status code %v", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status code %v", response.Status()))

	var endpoint Endpoints

	errJson := response.UnmarshalJson(&endpoint)

	if errJson != nil {
		return nil, errJson
	}

	fmt.Println(fmt.Sprintf("repository_url %v", endpoint.RepositoryUrl))

	return &endpoint, nil
}
