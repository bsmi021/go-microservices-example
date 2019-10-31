package domain

import (
	"fmt"
	"golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Brian", LastName: "Smith", Email: "brian@gmail.com"},
	}
	UserDao usersDaoInterface
)

func init(){
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

// UserDao common  for user data access object
type userDao struct{}

// GetUser returns a user with the ID specified
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
