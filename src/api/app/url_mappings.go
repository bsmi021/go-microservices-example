package app

import (
	"github.com/bsmi021/go-microservices-example/src/api/controllers/repositories"
)

func mapUrls() {
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}