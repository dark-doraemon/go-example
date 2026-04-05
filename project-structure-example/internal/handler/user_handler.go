package handler

import (
	service "project-structure/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	users := uh.service.GetAllUsers()
	ctx.JSON(200,users )
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	uh.service.GetUserByUUID()
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	uh.service.CreateUser()
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	uh.service.UpdateUser()
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	uh.service.DeleteUser()
}
