package main

import (
	"fmt"
	"go-backend/config"
	"go-backend/handler"
	"go-backend/repositories"
	"go-backend/routers"
	"go-backend/services"
	"net/http"
)

func main() {

	config.ConnectDb()

	repository := repositories.NewUserRepository()
	serviece := services.NewUserService(repository)
	handler := handler.NewUserHandler(serviece)

	routers.RegisterUserRoutes(handler)

	fmt.Println("Server Started")
	http.ListenAndServe(":8000", nil)
}
