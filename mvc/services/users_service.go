package services

import (
	"golang-microservices/mvc/utils"
	"golang-microservices/mvc/domain"
)

// GetUser returns a user from the database
func GetUser(userID int64) (*domain.User, *utils.ApplicationError){
	return domain.GetUser(userID)
}