package routes

import (
	"project-structure/internal/middleware"

	"github.com/gin-gonic/gin"
)

type IRoute interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...IRoute) {
	r.Use(middleware.AuthMiddleware())
	r.Use(middleware.RateLimitingMiddlware())

	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}
}
