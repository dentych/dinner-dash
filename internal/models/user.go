package models

import (
	"fmt"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	DisplayName  string    `json:"display_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type Family struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Members []User `json:"members"`
	Owner   *User   `json:"owner"`
}

func (f *Family) Validate() error {
	if f.Name == "" {
		return fmt.Errorf("family name required")
	}
	if f.Owner == nil {
		return fmt.Errorf("family owner required")
	}
	return nil
}

type Session struct {
	ID        int
	UserId    int
	SessionId string
	ValidTo   time.Time
}
