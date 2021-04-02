package api

import (
	"context"
	"gitlab.com/dentych/dinner-dash/internal/database"
	"gitlab.com/dentych/dinner-dash/internal/models"
	"time"
)

type UserApi struct {
	userRepo database.UserRepo
}

func NewUserApi(userRepo database.UserRepo) *UserApi {
	return &UserApi{userRepo: userRepo}
}

func (a *UserApi) GetById(ctx context.Context, userID string) (*models.User, error) {
	u, err := a.userRepo.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	return &models.User{
		ID:          u.ID,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		CreatedAt:   u.CreatedAt,
		FamilyID:    u.FamilyID,
	}, nil
}

func (a *UserApi) CreateUser(ctx context.Context, user models.User) error {
	dbUser := database.User{
		ID:          user.ID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		FamilyID:    nil,
		CreatedAt:   time.Now(),
	}
	err := a.userRepo.Insert(ctx, dbUser)
	return err
}
