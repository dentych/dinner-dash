package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gitlab.com/dentych/dinner-dash/internal/api"
	"net/http"
)

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

func Login(api *api.UserApi) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		var body map[string]string
		err := c.Bind(&body)
		if err != nil {
			c.Logger().Errorf("error binding request body: %s", err)
			return echo.ErrUnauthorized
		}

		username := body["username"]
		password := body["password"]

		token, err := api.Login(context.Background(), username, password)
		if err != nil {
			c.Logger().Errorf("failed to login: %s", err)
			return echo.ErrUnauthorized
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

func Register(api *api.UserApi) func(echo.Context) error {
	return func(c echo.Context) error {
		var body map[string]string
		err := c.Bind(&body)
		if err != nil {
			c.Logger().Errorf("Failed to bind user info: %s", err)
			return err
		}

		username := body["username"]
		password := body["password"]
		email := body["email"]

		if username == "" || password == "" || email == "" {
			return echo.ErrBadRequest
		}

		err = api.Register(context.Background(), username, password, email)
		if err != nil {
			c.Logger().Errorf("error registering new user: %s", err)
			br := echo.ErrBadRequest.SetInternal(err)
			return br
		}

		return nil
	}
}
