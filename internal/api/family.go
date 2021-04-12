package api

import (
	"context"
	"fmt"
	"gitlab.com/dentych/dinner-dash/internal/database"
	"gitlab.com/dentych/dinner-dash/internal/models"
	"math/rand"
	"time"
)

var (
	ErrUserNotInFamily = fmt.Errorf("user is not in the family")
)

type FamilyApi struct {
	familyRepo database.FamilyRepo
	userRepo   database.UserRepo
}

func NewFamilyApi(familyRepo database.FamilyRepo, userRepo database.UserRepo) *FamilyApi {
	return &FamilyApi{familyRepo: familyRepo, userRepo: userRepo}
}

func (api *FamilyApi) Create(ctx context.Context, userID string, familyName string) (int, error) {
	familyID, err := api.familyRepo.Insert(ctx, familyName, userID)
	if err != nil {
		return 0, err
	}
	return familyID, nil
}

func (api *FamilyApi) Get(ctx context.Context, userID string, familyID int) (models.Family, error) {
	userIsInFamily, err := api.familyRepo.UserInFamily(ctx, userID, familyID)
	if err != nil {
		return models.Family{}, err
	}

	if !userIsInFamily {
		return models.Family{}, ErrUserNotInFamily
	}

	dbFamily, err := api.familyRepo.GetById(ctx, familyID)
	if err != nil {
		return models.Family{}, err
	}

	members, err := api.userRepo.GetByFamilyId(ctx, familyID)
	if err != nil {
		return models.Family{}, err
	}

	var invitationID string
	if dbFamily.InvitationID != nil {
		invitationID = *dbFamily.InvitationID
	}

	family := models.Family{
		ID:           dbFamily.ID,
		Name:         *dbFamily.Name,
		InvitationId: invitationID,
	}
	for _, v := range members {
		family.Members = append(family.Members, models.FamilyMember{
			ID:          v.ID,
			DisplayName: *v.DisplayName,
		})
	}

	return family, nil
}

func (api *FamilyApi) Update(ctx context.Context, userID string, family models.UpdateFamilyInput) error {
	userInFamily, err := api.familyRepo.UserInFamily(ctx, userID, family.ID)
	if err != nil {
		return err
	}

	if !userInFamily {
		return ErrUserNotInFamily
	}

	var fieldsToUpdate []string
	if family.Name != nil {
		fieldsToUpdate = append(fieldsToUpdate, "name")
	}
	if family.InvitationID != nil {
		fieldsToUpdate = append(fieldsToUpdate, "invitationID")
	}
	err = api.familyRepo.Update(ctx, database.Family{
		ID:           family.ID,
		Name:         family.Name,
		InvitationID: family.InvitationID,
	}, fieldsToUpdate...)
	if err != nil {
		return err
	}

	return nil
}

func (api *FamilyApi) CreateInvitationLink(ctx context.Context, userID string, familyID int) (string, error) {
	invitationID := api.randomString(10)

	userInFamily, err := api.familyRepo.UserInFamily(ctx, userID, familyID)
	if err != nil {
		return "", err
	}

	if !userInFamily {
		return "", ErrUserNotInFamily
	}

	err = api.familyRepo.Update(ctx, database.Family{ID: familyID, InvitationID: &invitationID}, "InvitationID")

	if err != nil {
		return "", err
	}

	return invitationID, nil
}

func (api *FamilyApi) DeleteInvitationLink(ctx context.Context, userID string, familyID int) error {
	userInFamily, err := api.familyRepo.UserInFamily(ctx, userID, familyID)
	if err != nil {
		return err
	}

	if !userInFamily {
		return ErrUserNotInFamily
	}

	err = api.familyRepo.Update(ctx, database.Family{ID: familyID, InvitationID: nil}, "invitationID")
	if err != nil {
		return err
	}

	return nil
}

func (api *FamilyApi) AcceptInvitation(ctx context.Context, userID string, invitationID string) (int, error) {
	family, err := api.familyRepo.GetByInvitationID(ctx, invitationID)
	if err != nil {
		return 0, err
	}

	err = api.familyRepo.AddMember(ctx, family.ID, userID)
	if err != nil {
		return 0, err
	}

	return family.ID, nil
}

func (api *FamilyApi) LeaveFamily(ctx context.Context, userID string, familyID int) error {
	inFamily, err := api.familyRepo.UserInFamily(ctx, userID, familyID)
	if err != nil {
		return err
	}

	if !inFamily {
		return ErrUserNotInFamily
	}

	err = api.familyRepo.RemoveMember(ctx, familyID, userID)
	if err != nil {
		return err
	}

	return nil
}

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func (api *FamilyApi) randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	output := make([]rune, length)
	for i := range output {
		output[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(output)
}

func (api *FamilyApi) GetInvitationInformation(ctx context.Context, invitationID string) (models.InvitationInformation, error) {
	family, err := api.familyRepo.GetByInvitationID(ctx, invitationID)
	if err != nil {
		return models.InvitationInformation{}, err
	}

	return models.InvitationInformation{
		FamilyName:   *family.Name,
		InvitationID: *family.InvitationID,
	}, nil
}
