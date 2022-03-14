package transport

import (
	"REST-API-BookCatalog-Gin/entity"
)

type GeneralResponse struct {
	Message string `json:"message"`
}

type GetList struct {
	Count    int           `json:"count"`
	ListBook []entity.Book `json:"listBook"`
}

type GetBookResponse struct {
	Data entity.Book `json:"data"`
}

type ResponseError struct {
	Message string
	Status  int
}

// Get Users
type GetListUser struct {
	Count int `json:"count"`
	ListUser []entity.User `json:"listUser"`
}

type GetUserResponse struct {
	Data entity.User `json:"data"`
}