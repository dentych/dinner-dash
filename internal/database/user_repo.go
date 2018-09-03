package database

//go:generate mockgen -destination mocks/user_repo.go . UserRepo

import "gitlab.com/dentych/dinner-dash/internal/models"

type UserRepo interface {
	GetById(id int) (models.User, error)
	Insert(user models.User) (int, error)
}
