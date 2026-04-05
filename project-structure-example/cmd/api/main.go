package main

import (
	"project-structure/internal/config"
	"project-structure/internal/handler"
	repository "project-structure/internal/repositories"
	"project-structure/internal/routes"
	service "project-structure/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.NewConfig()

	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler)

	r := gin.Default()

	routes.RegisterRoutes(r, userRoutes)

	if err := r.Run(cfg.ServiceAddress); err != nil {
		panic(err)
	}

}
