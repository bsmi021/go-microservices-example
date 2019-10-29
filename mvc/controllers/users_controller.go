package controllers

import (
	"golang-microservices/mvc/utils"
	"encoding/json"
	"golang-microservices/mvc/services"
	"net/http"
	"strconv"
)

// GetUser returns a user from the underlying service
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userID, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id_must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		// Just return the Bad Request to the client
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)

	if apiErr != nil {
		// handle error and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)

	resp.Write(jsonValue)

}
