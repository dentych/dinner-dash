package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gitlab.com/dentych/dinner-dash/internal/api"
	"gitlab.com/dentych/dinner-dash/internal/models"
	"net/http"
	"time"
)

type createUserInput struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

func GetUser(userApi *api.UserApi) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		sub := claims["sub"]
		userID := sub.(string)

		userInfo, err := userApi.GetById(context.Background(), userID)
		if err != nil {
			c.Logger().Errorf("getuser api returned error: %s", err)
			return c.JSON(http.StatusInternalServerError, "there was an error")
		}
		if userInfo == nil {
			return c.JSON(http.StatusNotFound, "user not found in database")
		}
		return c.JSON(http.StatusOK, userInfo)
	}
}

func CreateUser(userApi *api.UserApi) func(c echo.Context) error {
	return func(c echo.Context) error {
		userClaims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
		userID := userClaims["sub"].(string)
		var input createUserInput
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Errorf("error binding: %s", err)
			return c.JSON(http.StatusBadRequest, "invalid input")
		}

		user := models.User{
			ID:          userID,
			Email:       input.Email,
			DisplayName: &input.DisplayName,
			CreatedAt:   time.Time{},
			FamilyID:    nil,
		}
		err = userApi.CreateUser(context.Background(), user)
		if err != nil {
			c.Logger().Errorf("error creating user: %s", err)
			return c.JSON(http.StatusInternalServerError, "internal error")
		}
		return c.JSON(http.StatusCreated, user)
	}
}
