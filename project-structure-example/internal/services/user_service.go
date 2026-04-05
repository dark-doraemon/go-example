package service

import (
	"project-structure/internal/models"
	repository "project-structure/internal/repositories"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetAllUsers() []models.User{
	return us.repo.FindAll()
}

func (us *UserService) GetUserByUUID() {

}

func (us *UserService) CreateUser() {

}

func (us *UserService) UpdateUser() {

}

func (us *UserService) DeleteUser() {

}
