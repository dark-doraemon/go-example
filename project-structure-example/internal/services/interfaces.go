package service

import "project-structure/internal/models"

type IUserService interface {
	GetAllUsers() []models.User
	GetUserByUUID()
	CreateUser()
	UpdateUser()
	DeleteUser()
}
