package github_provider

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/bsmi021/go-microservices-example/src/api/clients/rest_client"
	"net/http"
	"fmt"
	"github.com/bsmi021/go-microservices-example/src/api/domain/github"
)

const(
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

// CreateRepo will attempt to create a new repository in the provided account's GitHub account
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse){
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	// execute a post on the provided URI for creating a new repo
	response, err := rest_client.Post(urlCreateRepo, request, headers)
	if err != nil {
		// if an error occurs log the output and return the errorresponse with the status code and message
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}

	// read the contents from the response.Body
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// on error return the formatted error
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	// deallocate the body stream when the function exits, even if there is an error
	defer response.Body.Close()
	
	// if the status code is above 299 an error occured with the payload
	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	// got here so create the response by unloading the JSON into a CreateRepoResponse
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		//if there was an error dump an error
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal response"}
	}

	// good case return the result
	return &result, nil
	
}