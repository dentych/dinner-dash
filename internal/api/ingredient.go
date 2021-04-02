package api

import (
	"gitlab.com/dentych/dinner-dash/internal/database"
	"gitlab.com/dentych/dinner-dash/internal/models"
)

type IngredientApi struct {
	ingredientRepo database.IngredientRepo
}

func NewIngredientApi(ingredientRepo database.IngredientRepo) *IngredientApi {
	return &IngredientApi{ingredientRepo: ingredientRepo}
}

func (api *IngredientApi) Create(user models.User, ingredient models.Ingredient) (int, error) {
	return -1, nil
}