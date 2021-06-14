package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/dentych/dinner-dash/internal/api"
	"gitlab.com/dentych/dinner-dash/internal/models"
)

type newRecipe struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Ingredients []models.IngredientInput `json:"ingredients"`
}

func GetRecipes(c echo.Context) error {
	return c.JSON(200, []interface{}{})
}

func AddRecipe(api *api.RecipeApi) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var recipe newRecipe
		err := ctx.Bind(&recipe)
		if err != nil {
			ctx.Logger().Infof("User tried to create a new recipe which was invalid: %s", err)
			return ctx.JSON(400, "invalid recipe")
		}
		return fmt.Errorf("not implemented yet")
	}
}

func mapIngredients(ingredients []models.IngredientInput) []models.Ingredient {
	//var outputIngredients []models.Ingredient
	//for _, ingredient := range ingredients {
	//	output := models.Ingredient{
	//		Type:   models.IngredientType{},
	//		Amount: 0,
	//	}
	//}
	return nil
}
