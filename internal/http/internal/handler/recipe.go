package handler

import "github.com/labstack/echo/v4"

func GetRecipes(c echo.Context) error {
	return c.JSON(200, []interface{}{})
}
