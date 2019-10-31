package app

import (
	"golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}