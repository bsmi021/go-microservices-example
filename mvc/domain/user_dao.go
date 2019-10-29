package domain

import (
	"net/http"
	"golang-microservices/mvc/utils"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Brian", LastName: "Smith", Email: "brian@gmail.com"},
	}
)

// GetUser returns a user with the ID specified
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}
	
}
