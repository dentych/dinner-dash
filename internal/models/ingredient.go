package models

type Ingredient struct {
	ID     int            `json:"id"`
	Type   IngredientType `json:"type"`
	Amount float64        `json:"amount"`
}

type IngredientType struct {
	// ID of the ingredient in the database
	ID int `json:"id"`
	// Name of the ingredient
	Name string `json:"name"`
	// Unit is the SI Unit to measure this ingredient in
	Unit string `json:"unit"`
}

type IngredientInput struct {
	ID               int     `json:"id"`
	IngredientTypeID int     `json:"ingredient_type_id"`
	Amount           float64 `json:"amount"`
}
