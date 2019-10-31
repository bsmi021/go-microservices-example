package services

import (
	"golang-microservices/mvc/utils"
	"golang-microservices/mvc/domain"
)

type usersService struct {
}

var (
	UsersService usersService
)

// GetUser returns a user from the database
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError){
	return domain.UserDao.GetUser(userID)
}