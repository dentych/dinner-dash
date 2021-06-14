package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gitlab.com/dentych/dinner-dash/internal/database"
	"gitlab.com/dentych/dinner-dash/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
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
		Username:   u.Username,
		Email:      u.Email,
		CreatedAt:  u.CreatedAt,
		FamilyID:   u.FamilyID,
		FamilyName: u.FamilyName,
	}, nil
}

func (a *UserApi) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := a.userRepo.GetById(ctx, username)
	if err != nil {
		return "", echo.ErrUnauthorized
	}

	if user == nil {
		return "", echo.ErrUnauthorized
	}

	if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)) != nil {
		return "", echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("supersecretkey"))
	if err != nil {
		log.Printf("failed to sign token: %s", err)
		return "", err
	}

	return t, nil
}

func (a *UserApi) Register(ctx context.Context, username string, password string, email string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = a.userRepo.Insert(ctx, database.User{
		Username:       username,
		Email:          email,
		HashedPassword: string(hashedPass),
		CreatedAt:      time.Now(),
	})

	return err
}
