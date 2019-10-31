package services

import (
	"golang-microservices/src/api/log/option_b"
	"net/http"
	"sync"
	"golang-microservices/src/api/config"
	"golang-microservices/src/api/providers/github_provider"
	"golang-microservices/src/api/domain/github"
	"strings"
	"golang-microservices/src/api/utils/errors"
	
	"golang-microservices/src/api/domain/repositories"
	
)

type reposService struct {}

type reposServiceInterface interface {
	CreateRepo(clientID string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	// sets the value of the variable
	RepositoryService = &reposService{}
}

// CreateRepo requires a generic CreateRepoRequest which ensures that a Name and Description are provided for the new source control repository
func (s *reposService) CreateRepo(clientID string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError){
	input.Name = strings.TrimSpace(input.Name)
	
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// since only github is available it makes sense to only have a github request be created
	// the repositories.CreateRepoRequest can be expanded to chose from different optoins
	request := github.CreateRepoRequest{
		Name: input.Name,
		Private: false,
		Description: input.Description,
	}
	// log request
	option_b.Info("about to send request to external api",
		option_b.Field("client_id", clientID),
		option_b.Field("status", "pending"),
		option_b.Field("authenticated", clientID != ""))

	// like the request, the response is shaped agains the github_provider.CreateRepo request, which requires the access token (from secrets)
	// and a github.CreateRepoRequest, note that the generic repositories.CreateRepoRequest can be augmented to include a repo type
	// when more providers are created
	response, err := github_provider.CreateRepo(config.GetGitHubAccessToken(), request)
	if err != nil {
		// log error
		option_b.Error("response obtained from external api", err,
						option_b.Field("client_id", clientID),
						option_b.Field("status", "error"),
						option_b.Field("authenticated", clientID != ""))
		// gets back a github.ErrorResponse create a new apiErr, again like above this is shaped to github request/response/error
		// if more providers are created later this can be extended or redesigned
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	// log response
	option_b.Info("response obtained from external api",
			option_b.Field("client_id", clientID),
			option_b.Field("status", "success"),
			option_b.Field("authenticated", clientID != ""))

	// shape the response to the generic repositories.CreateRepoResponse and then return it out
	result := repositories.CreateRepoResponse {
		ID: response.ID,
		Name: response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil

}

func (s *reposService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError){
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0

	for _, current := range result.Results {
		if current.Response != nil {
			successCreations ++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error: incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	result, err := s.CreateRepo("TODO_client_id" , input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}