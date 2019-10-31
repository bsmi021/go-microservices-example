package services

import (
	"net/http"
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)
// GetItem returns an item based on the provided ID value
func (i *itemsService) GetItem(itemID string)(*domain.Item, *utils.ApplicationError){
	return nil, &utils.ApplicationError {
		Message: "not implemented",
		StatusCode: http.StatusBadRequest,
		Code: "bad_request",
	}
}