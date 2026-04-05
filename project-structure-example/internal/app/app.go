package app

import (
	"project-structure/internal/config"
	"project-structure/internal/routes"

	"github.com/gin-gonic/gin"
)

type IModule interface {
	Route() routes.IRoute
}

type Application struct {
	config *config.Config
	router *gin.Engine
}

func NewApplication(cfg *config.Config) *Application {

	r := gin.Default()

	modules := []IModule{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
	}
}

func (app *Application) Run() error {
	return app.router.Run(app.config.ServiceAddress)
}

func getModuleRoutes(modules []IModule) []routes.IRoute {
	routelist := make([]routes.IRoute, len(modules))

	for i := 0; i < len(modules); i++ {
		routelist[i] = modules[i].Route()
	}

	return routelist
}
