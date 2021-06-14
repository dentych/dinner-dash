package models

import (
	"time"
)

type User struct {
	Username   string    `json:"id"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	FamilyID   *int      `json:"familyId"`
	FamilyName *string   `json:"familyName"`
}

type UserWithFamilyName struct {
	Username   string    `json:"id"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	FamilyID   *int      `json:"familyId,omitempty"`
	FamilyName *string   `json:"familyName,omitempty"`
}

type FamilyMember struct {
	Username string `json:"displayName"`
}

type Family struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	InvitationId string         `json:"invitationId,omitempty"`
	Members      []FamilyMember `json:"members"`
}

type UpdateFamilyInput struct {
	ID           int
	Name         *string
	InvitationID *string
}

type Session struct {
	ID        int
	UserId    int
	SessionId string
	ValidTo   time.Time
}
