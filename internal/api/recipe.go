package api

import (
	"gitlab.com/dentych/dinner-dash/internal/database"
)

type RecipeApi struct {
	familyRepo database.FamilyRepo
	recipeRepo database.RecipeRepo
}

func NewRecipeApi(familyRepo database.FamilyRepo, recipeRepo database.RecipeRepo) *RecipeApi {
	return &RecipeApi{familyRepo: familyRepo, recipeRepo: recipeRepo}
}

//func (api *RecipeApi) Create(user models.User, familyID int, recipe models.RecipeInput) (int, error) {
//	userInFamily, err := api.familyRepo.UserInFamily(user.ID, familyID)
//	if err != nil {
//		return -1, err
//	}
//	if !userInFamily {
//		return -1, ErrUserNotInFamily
//	}
//	validationErr := recipe.Validate()
//	if validationErr != nil {
//		return -1, err
//	}
//	recipeID, err := api.recipeRepo.Insert(recipe)
//	if err != nil {
//		return -1, err
//	}
//	return recipeID, nil
//}

//func (api *RecipeApi) GetRecipesInFamily(user models.User, familyID int) ([]models.Recipe, error) {
//	userInFamily, err := api.familyRepo.UserInFamily(user.ID, familyID)
//	if err != nil {
//		return []models.Recipe{}, err
//	}
//	if !userInFamily {
//		return []models.Recipe{}, ErrUserNotInFamily
//	}
//	recipes, err := api.recipeRepo.GetByFamilyID(familyID)
//	if err != nil {
//		return []models.Recipe{}, err
//	}
//
//	return recipes, nil
//}

//func (api *RecipeApi) UpdateRecipe(user models.User, familyID int, recipe models.RecipeInput) error {
//	userInFamily, err := api.familyRepo.UserInFamily(user.ID, familyID)
//	if err != nil {
//		return nil
//	}
//	if !userInFamily {
//		return ErrUserNotInFamily
//	}
//	validationErr := recipe.Validate()
//	if validationErr != nil {
//		return validationErr
//	}
//	return api.recipeRepo.Update(recipe)
//}
