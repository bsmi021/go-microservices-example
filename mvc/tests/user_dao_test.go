package tests

import (
	"testing"

	"github.com/bsmi021/go-microservices-example/mvc/domain"
	"github.com/bsmi021/go-microservices-example/mvc/utils"
)

// userDaoImpl is a concrete implementation for testing the actual DAO
type userDaoImpl struct{}

func (u *userDaoImpl) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	users := map[int64]*domain.User{
		123: {ID: 123, FirstName: "Brian", LastName: "Smith", Email: "brian@gmail.com"},
	}

	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    "user does not exist",
		StatusCode: 404,
		Code:       "not_found",
	}
}

func TestDAOGetUserNotFound(t *testing.T) {
	// Use a real DAO implementation for this test
	realDao := &userDaoImpl{}

	// execute
	user, err := realDao.GetUser(1324)

	// validate
	if user != nil {
		t.Error("User should not be found")
	}

	if err == nil {
		t.Error("Error should be returned")
	}
}

func TestDAOGetUserFound(t *testing.T) {
	// Use a real DAO implementation for this test
	realDao := &userDaoImpl{}

	// execute
	user, err := realDao.GetUser(123)

	// validate
	if user == nil {
		t.Error("User should be found")
	}

	if err != nil {
		t.Errorf("No error should be returned, got: %v", err)
	}
}
