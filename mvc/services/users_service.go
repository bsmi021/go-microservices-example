package services

import (
	"github.com/bsmi021/go-microservices-example/mvc/utils"
	"github.com/bsmi021/go-microservices-example/mvc/domain"
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