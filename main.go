package main

import (
	"REST-API-BookCatalog-Gin/config"
	"REST-API-BookCatalog-Gin/server"

	"github.com/go-playground/validator/v10"
)

func main() {
	validation := validator.New()
	cfg := config.LoadConfig()
	dbInit, err := config.MySQL(cfg)
	if err != nil {
		panic(err)
	}

	server := server.NewServer(dbInit, validation)

	server.ListenAndServer("8080")
}
