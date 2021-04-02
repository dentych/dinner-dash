package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetUserID(ctx echo.Context) string {
	return ctx.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
}
