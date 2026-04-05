package repository

import "project-structure/internal/models"

type IUserRepository interface {
	FindAll() []models.User
	GetUserByUUID()
	CreateUser()
	UpdateUser()
	DeleteUser()
}
