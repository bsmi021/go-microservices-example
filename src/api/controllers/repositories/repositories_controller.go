package repositories

// uses gin.gonic api

import (
	"github.com/bsmi021/go-microservices-example/src/api/domain/repositories"
	"github.com/bsmi021/go-microservices-example/src/api/utils/errors"
	"github.com/bsmi021/go-microservices-example/src/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	// create a variable for the CreateRepoRequest
	var request repositories.CreateRepoRequest

	// attempt to assign to the request variable with JSON by checking to see 
	// if it's valid json and can be used to populate CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.BadRequestAPIError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	clientID := c.GetHeader("X-Client-Id")

	result, err := services.RepositoryService.CreateRepo(clientID, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.BadRequestAPIError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(result.StatusCode, result)
}