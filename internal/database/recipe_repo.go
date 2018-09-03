package database

//go:generate mockgen -destination mocks/recipe_repo.go . RecipeRepo

import "gitlab.com/dentych/dinner-dash/internal/models"

type RecipeRepo interface {
	GetById(id int) (models.Recipe, error)
	GetByFamilyID(familyID int) ([]models.Recipe, error)
	Insert(recipe models.RecipeInput) (int, error)
	Update(recipe models.RecipeInput) error
}
