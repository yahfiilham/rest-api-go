package usecase

import (
	// "REST-API-BookCatalog-Gin/entity"
	"REST-API-BookCatalog-Gin/entity"
	"REST-API-BookCatalog-Gin/repository"
	"REST-API-BookCatalog-Gin/transport"

	"database/sql"
	"net/http"
)

type UserUsecase interface {
	GetListUser() (*transport.GetListUser, *transport.ResponseError)
	GetUserByID(id int) (*transport.GetUserResponse, *transport.ResponseError)
	AddUser(data transport.InsertUser) (*transport.GeneralResponse, *transport.ResponseError)
	UpdateUser(id int, data transport.UpdateUser) (*transport.GeneralResponse, *transport.ResponseError)
	// DeleteUser(id int) (*transport.GetUserResponse, *transport.ResponseError)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		repository: userRepository,
	}
}

func (u *userUsecase) GetListUser() (*transport.GetListUser, *transport.ResponseError) {
	result, err := u.repository.GetUsers()

	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	return &transport.GetListUser{
		Count:    len(result),
		ListUser: result,
	}, nil
}

func (u *userUsecase) GetUserByID(id int) (*transport.GetUserResponse, *transport.ResponseError) {
	result, err := u.repository.GetUserByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			responseError := &transport.ResponseError{
				Message: "Data not found, please check your request.",
				Status:  http.StatusNotFound,
			}
			return nil, responseError
		}
	}

	return &transport.GetUserResponse{
		Data: *result,
	}, nil
}

func (u *userUsecase) AddUser(data transport.InsertUser) (*transport.GeneralResponse, *transport.ResponseError) {
	createPayload := entity.User{
		Username:    data.Username,
		Email: data.Email,
	}

	err := u.repository.AddUser(createPayload)
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	result := &transport.GeneralResponse{
		Message: "Success to insert book : " + createPayload.Username + " into database",
	}

	return result, nil
}

func (u *userUsecase) UpdateUser(id int, data transport.UpdateUser) (*transport.GeneralResponse, *transport.ResponseError) {
	result, err := u.repository.GetUserByID(id)

	if err != nil {
		responseError := &transport.ResponseError{
			Message: "Data not found, please check your request.",
			Status:  http.StatusNotFound,
		}
		return nil, responseError
	}

	if data.Username == "" {
		data.Username = result.Username
	}

	if data.Email == "" {
		data.Email = result.Email
	}

	createPayload := entity.User{
		Id:      id,
		Username:    data.Username,
		Email: data.Email,
	}

	err = u.repository.UpdateUser(createPayload)
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	res := &transport.GeneralResponse{
		Message: "Success to update book : " + createPayload.Username + ".",
	}

	return res, nil
}

/*
func (u *userUsecase) DeleteUser(id int) (*transport.GeneralResponse, *transport.ResponseError) {
	_, err := u.repository.GetUserByID(id)

	if err != nil {
		responseError := &transport.ResponseError{
			Message: "Data not found, please check your request.",
			Status:  http.StatusNotFound,
		}
		return nil, responseError
	}

	err = u.repository.DeleteUser(id)
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	res := &transport.GeneralResponse{
		Message: "Success to deleted!",
	}

	return res, nil
}
*/
