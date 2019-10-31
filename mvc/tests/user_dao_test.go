package tests

import (

	"golang-microservices/mvc/domain"
	"testing"
)



func TestDAOGetUserNotFound(t *testing.T){
	// execute
	if user, _ := domain.UserDao.GetUser(1324); user == nil {
		t.Log(user)
		return
	}

	// validate
	t.Error("User was found")
}

func TestDAOGetUserFound(t *testing.T){
		// execute
	if user, _ := domain.UserDao.GetUser(123); user != nil {
		t.Log(user)
		return
	}

	// validate
	t.Error("User not found")
}
