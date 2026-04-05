package app

import (
	"project-structure/internal/handler"
	repository "project-structure/internal/repositories"
	"project-structure/internal/routes"
	service "project-structure/internal/services"
)

type UserModule struct {
	route routes.IRoute
}

func NewUserModule() *UserModule {
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler)

	return &UserModule{
		route: userRoutes,
	}
}

func (m *UserModule) Route() routes.IRoute {
	return m.route
}
