package server

import (
	"REST-API-BookCatalog-Gin/handler"
	"REST-API-BookCatalog-Gin/repository"
	"REST-API-BookCatalog-Gin/usecase"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiServer struct {
	DB        *sql.DB
	Router    *gin.Engine
	validator *validator.Validate
}

func NewServer(db *sql.DB, validator *validator.Validate) *ApiServer {
	r := gin.New()
	return &ApiServer{
		DB:        db,
		Router:    r,
		validator: validator,
	}
}

func (server *ApiServer) ListenAndServer(port string) {
	server.Router.Use(gin.Logger())
	server.registerRouter()

	http.ListenAndServe(":"+port, server.Router)
}

func (server *ApiServer) registerRouter() {
	repo := repository.NewBookRepository(server.DB)
	uCase := usecase.NewBookUsecase(repo)
	bookHandler := handler.NewBookHandler(uCase, server.validator)

	userRepo := repository.NewUserRepository(server.DB)
	userCase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userCase, server.validator)


	server.Router.GET("api/books", bookHandler.GetList)
	server.Router.GET("api/books/:bookID", bookHandler.GetByID)
	server.Router.POST("api/insertbook", bookHandler.AddBook)
	server.Router.PATCH("api/updatebook/:bookID", bookHandler.UpdateBook)
	
	server.Router.GET("api/users", userHandler.GetListUser)
	server.Router.GET("api/users/:userID", userHandler.GetUserByID)
	server.Router.POST("api/users", userHandler.AddUser)
	server.Router.PATCH("api/users/:userID", userHandler.UpdateUser)
}
