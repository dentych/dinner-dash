package handler

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gitlab.com/dentych/dinner-dash/internal/api"
	"gitlab.com/dentych/dinner-dash/internal/http/internal/util"
	"gitlab.com/dentych/dinner-dash/internal/models"
	"net/http"
	"strconv"
)

type createFamilyInput struct {
	Name string `json:"name"`
}

type editFamilyInput struct {
	ID   int     `json:"id"`
	Name *string `json:"name,omitempty"`
}

func GetFamily(familyApi *api.FamilyApi) func(echo.Context) error {
	return func(c echo.Context) error {
		familyId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Errorf("error converting family ID param to int: %s", err)
			return c.JSON(http.StatusBadRequest, "invalid id - must be an integer")
		}
		claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)
		family, err := familyApi.Get(context.Background(), userID, familyId)
		if err != nil {
			if errors.Is(err, api.ErrUserNotInFamily) {
				return c.JSON(http.StatusForbidden, "user not in family")
			}
			c.Logger().Errorf("error getting family: %s", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, family)
	}
}

func CreateFamily(familyApi *api.FamilyApi) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		userID := util.GetUserID(ctx)
		var input createFamilyInput
		err := ctx.Bind(&input)
		if err != nil {
			ctx.Logger().Errorf("Error binding create family input: %s", err)
			return ctx.JSON(http.StatusBadRequest, "invalid input")
		}
		familyID, err := familyApi.Create(context.Background(), userID, input.Name)
		if err != nil {
			ctx.Logger().Errorf("Error creating family: %s", err)
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusCreated, familyID)
	}
}

func UpdateFamily(familyApi *api.FamilyApi) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		userID := util.GetUserID(ctx)
		var input editFamilyInput
		err := ctx.Bind(&input)
		if err != nil {
			ctx.Logger().Errorf("error binding family input, userID=%s:", userID, err)
			return ctx.JSON(http.StatusBadRequest, "invalid input")
		}

		err = familyApi.Update(context.Background(), userID, models.UpdateFamilyInput{ID: input.ID, Name: input.Name})
		if err != nil {
			if errors.Is(err, api.ErrUserNotInFamily) {
				return ctx.JSON(http.StatusForbidden, "user not in family")
			}
			ctx.Logger().Errorf("error updating family familyID=%d: %s", input.ID, err)
			return ctx.JSON(http.StatusInternalServerError, "failed to update family")
		}
		return ctx.JSON(200, "family updated")
	}
}

func GenerateInvitation(familyApi *api.FamilyApi) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		userID := util.GetUserID(ctx)
		idParam := ctx.Param("id")
		familyID, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.Logger().Errorf("failed to bind familyID=%s: %s", idParam, err)
			return ctx.JSON(http.StatusBadRequest, "invalid family ID: "+idParam)
		}
		invitationID, err := familyApi.CreateInvitation(context.Background(), userID, familyID)
		if err != nil {
			ctx.Logger().Errorf("error creating invitation ID for familyID=%d, userID=%s: %s", familyID, userID, err)
			return ctx.JSON(http.StatusInternalServerError, "failed to create invitation ID")
		}
		return ctx.JSON(200, invitationID)
	}
}

func AcceptInvitation(familyApi *api.FamilyApi) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		userID := util.GetUserID(ctx)
		invitationID := ctx.Param("invitationId")
		if invitationID == "" {
			ctx.Logger().Errorf("invalid invitationID: %s", invitationID)
			return ctx.JSON(http.StatusBadRequest, "invalid invitationID")
		}
		familyID, err := familyApi.AcceptInvitation(context.Background(), userID, invitationID)
		if err != nil {
			ctx.Logger().Errorf("failed to accept invitation with ID=%s for userID=%s and familyID=%s. error: %s", invitationID, userID, familyID, err)
			return ctx.JSON(http.StatusInternalServerError, "failed to accept invitation")
		}

		return ctx.JSON(http.StatusOK, familyID)
	}
}
