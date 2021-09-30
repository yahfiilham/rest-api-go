package handler

import (
	"REST-API-BookCatalog-Gin/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (b *bookHandler) GetList(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")
	result, err := b.usecase.GetList()

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (b *bookHandler) GetByID(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")
	log.Println("ididididid")
	id, _ := strconv.Atoi(gc.Param("bookID"))

	log.Println(id)

	result, err := b.usecase.GetByID(id)
	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}
