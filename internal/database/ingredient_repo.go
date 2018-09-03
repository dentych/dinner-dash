package database

//go:generate mockgen -destination mocks/ingredient_repo.go . IngredientRepo

import "gitlab.com/dentych/dinner-dash/internal/models"

type IngredientRepo interface {
	Insert(ingredient models.IngredientInput) (int, error)
	GetAll() ([]models.Ingredient, error)
	Update(ingredient models.Ingredient) error
	Delete(ingredient models.Ingredient) error
}
