package models

import "fmt"

type Recipe struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Url         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	Family      *Family `json:"family"`
	Ingredients []Ingredient
}

func (r *Recipe) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("recipe name required")
	}
	if r.Family == nil {
		return fmt.Errorf("recipe must be part of a family")
	}
	return nil
}

type RecipeInput struct {
	Name        string            `json:"name"`
	Url         *string           `json:"url,omitempty"`
	Description *string           `json:"description,omitempty"`
	Directions  *string           `json:"directions,omitempty"`
	FamilyID    int               `json:"family_id"`
	Ingredients []IngredientInput `json:"ingredients"`
}

func (r *RecipeInput) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("recipe name required")
	}
	if r.FamilyID == 0 {
		return fmt.Errorf("family_id required")
	}
	return nil
}
