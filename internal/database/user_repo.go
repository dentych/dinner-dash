package database

//go:generate mockgen -destination mocks/user_repo.go . UserRepo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserRepo interface {
	GetById(ctx context.Context, userID string) (*User, error)
	GetByFamilyId(ctx context.Context, familyID int) ([]User, error)
	Insert(ctx context.Context, user User) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

type User struct {
	ID          string
	Email       string
	DisplayName *string   `db:"display_name"`
	FamilyID    *int      `db:"family_id"`
	CreatedAt   time.Time `db:"created_at"`
}

func (r *userRepo) GetById(ctx context.Context, userID string) (*User, error) {
	var user User
	err := r.db.Get(&user, "select * from public.user where id = $1", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByFamilyId(ctx context.Context, familyID int) ([]User, error) {
	var users []User
	err := r.db.Select(&users, "select * from public.user where family_id = $1", familyID)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Insert(ctx context.Context, user User) error {
	_, err := r.db.Exec("insert into public.user (id, email, display_name) values ($1, $2, $3)", user.ID, user.Email, user.DisplayName)
	if err != nil {
		return err
	}
	return nil
}
