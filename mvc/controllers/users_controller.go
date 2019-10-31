package controllers

import (
	"log"
	"github.com/gin-gonic/gin"
	"golang-microservices/mvc/utils"
	"golang-microservices/mvc/services"
	"net/http"
	"strconv"
)

// GetUser returns a user from the underlying service
func GetUser(c *gin.Context) {

	userID, err := (strconv.ParseInt(c.Param("user_id"), 10, 64))

	if err != nil {
		log.Println(err)
		apiErr := &utils.ApplicationError{
			Message: "user_id_must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	//item, itemErr := services.ItemsService.GetItem("none")

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	// return user to client
	utils.Respond(c, http.StatusOK, user)

}
