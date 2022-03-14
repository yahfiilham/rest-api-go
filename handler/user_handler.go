package handler

import (
	// "REST-API-BookCatalog-Gin/transport"
	"REST-API-BookCatalog-Gin/transport"
	"REST-API-BookCatalog-Gin/usecase"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	usecase   usecase.UserUsecase
	validator *validator.Validate
}

func NewUserHandler(userUsecase usecase.UserUsecase, validator *validator.Validate) *userHandler {
	return &userHandler{
		usecase:   userUsecase,
		validator: validator,
	}
}

func (u *userHandler) GetListUser(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")
	result, err := u.usecase.GetListUser()

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (u *userHandler) GetUserByID(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(gc.Param("userID"))

	result, err := u.usecase.GetUserByID(id)
	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (u *userHandler) AddUser(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")

	var requestUser transport.InsertUser
	if err := gc.ShouldBindJSON(&requestUser); err != nil {
		gc.JSON(http.StatusBadRequest, transport.ResponseError{
			Message: "error while decode request body",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// checking validation
	errorValidation := u.validator.Struct(requestUser)
	if errorValidation != nil {
		gc.JSON(http.StatusBadRequest, transport.ResponseError{
			Message: errorValidation.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	result, responseError := u.usecase.AddUser(requestUser)
	if responseError != nil {
		gc.JSON(responseError.Status, responseError)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (u *userHandler) UpdateUser(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(gc.Param("userID"))

	var requestUser transport.UpdateUser
	if err := gc.ShouldBindJSON(&requestUser); err != nil {
		gc.JSON(http.StatusBadRequest, transport.ResponseError{
			Message: "error while decode request body",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// checking validation
	errorValidation := u.validator.Struct(requestUser)
	if errorValidation != nil {
		gc.JSON(http.StatusBadRequest, transport.ResponseError{
			Message: errorValidation.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	result, responseError := u.usecase.UpdateUser(id, requestUser)
	if responseError != nil {
		gc.JSON(responseError.Status, responseError)
		return
	}

	gc.JSON(http.StatusOK, result)
}
