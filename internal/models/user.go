package models

import (
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	DisplayName *string   `json:"displayName"`
	CreatedAt   time.Time `json:"createdAt"`
	FamilyID    *int      `json:"familyId"`
}

type FamilyMember struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
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
