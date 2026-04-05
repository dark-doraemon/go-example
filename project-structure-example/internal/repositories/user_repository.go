package repository

import (
	"project-structure/internal/models"
)

type InMemoryUserRepository struct {
	repo []models.User
}

func NewInMemoryUserRepository() IUserRepository {
	return &InMemoryUserRepository{
		repo: make([]models.User, 0),
	}
}

func (repo *InMemoryUserRepository) FindAll() []models.User {
	return repo.repo
}

func (repo *InMemoryUserRepository) GetUserByUUID() {
}

func (repo *InMemoryUserRepository) CreateUser() {

}
func (repo *InMemoryUserRepository) UpdateUser() {

}
func (repo *InMemoryUserRepository) DeleteUser() {

}
