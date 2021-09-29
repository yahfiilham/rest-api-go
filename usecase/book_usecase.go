package usecase

import (
	"REST-API-BookCatalog-Gin/repository"
)

type BookUsecase interface {
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) BookUsecase {
	return &bookUsecase{
		repository: bookRepository,
	}
}
