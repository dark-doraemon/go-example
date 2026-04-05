package routes

import (
	"project-structure/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) IRoute {
	return &UserRoute{
		handler: handler,
	}
}

func (ur *UserRoute) Register(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", ur.handler.GetAllUsers)
		users.POST("", ur.handler.CreateUser)
		users.GET("/:uuid", ur.handler.GetUserByUUID)
		users.DELETE("/:uuid", ur.handler.DeleteUser)
		users.PUT("/:uuid", ur.handler.UpdateUser)
	}
}
