package api

import (
	"gitlab.com/dentych/dinner-dash/internal/database"
	"gitlab.com/dentych/dinner-dash/internal/models"
)

type RecipeApi struct {
	familyRepo database.FamilyRepo
	recipeRepo database.RecipeRepo
}

func NewRecipeApi(familyRepo database.FamilyRepo, recipeRepo database.RecipeRepo) *RecipeApi {
	return &RecipeApi{familyRepo: familyRepo, recipeRepo: recipeRepo}
}

func (a *RecipeApi) AddRecipe(recipe models.Recipe) error {
	return nil
}

