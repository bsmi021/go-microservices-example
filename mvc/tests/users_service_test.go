package tests

import (
	"github.com/bsmi021/go-microservices-example/mvc/domain"
	"github.com/bsmi021/go-microservices-example/mvc/utils"
	"github.com/bsmi021/go-microservices-example/mvc/services"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {}

func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError){
	return getUserFunction(userID)
}

func TestServiceGetUserFound (t *testing.T){
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError){
		return &domain.User{
			ID: 123, 
			FirstName: "Brian", 
			LastName: "Smith", 
			Email: "god@godsstuf.com",
		}, nil
	}
	usr, err := services.UsersService.GetUser(123)

	assert.NotNil(t, usr, "User found")
	assert.Nil(t, err, "Value nil")
	assert.EqualValues(t, usr.ID, 123, "Correct value")
	assert.EqualValuesf(t, usr.FirstName, "Brian", "Correct value")
	assert.EqualValues(t, usr.LastName, "Smith", "Correct value")
}

func TestServiceGetUserNotFound (t *testing.T){
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError){
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "user 0 does not exist",
			Code: "not_found",
		}
	}
	
	usr, err := services.UsersService.GetUser(0)

	assert.Nil(t, usr, "User is nil")
	assert.NotNil(t, err, "Error returned")

	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Code, "not_found")

}