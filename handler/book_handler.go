package handler

import (
	"REST-API-BookCatalog-Gin/usecase"

	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	usecase   usecase.BookUsecase
	validator *validator.Validate
}

func NewBookHandler(bookUseCase usecase.BookUsecase, validator *validator.Validate) *bookHandler {
	return &bookHandler{
		usecase:   bookUseCase,
		validator: validator,
	}
}
