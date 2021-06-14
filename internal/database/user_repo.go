package database

//go:generate mockgen -destination mocks/user_repo.go . UserRepo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	Username       string
	Email          string
	HashedPassword string    `db:"hashed_password"`
	FamilyID       *int      `db:"family_id"`
	FamilyName     *string   `db:"family_name"`
	CreatedAt      time.Time `db:"created_at"`
}

type UserRepo interface {
	GetById(ctx context.Context, username string) (*User, error)
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

func (r *userRepo) GetById(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.db.Get(&user, "select u.*, f.name family_name from public.user u left join public.family f on u.family_id = f.id where u.username = $1", username)
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
	sqlStatement := "insert into public.user (username, email, hashed_password, created_at) values ($1, $2, $3, $4)"
	_, err := r.db.Exec(sqlStatement, user.Username, user.Email, user.HashedPassword, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
