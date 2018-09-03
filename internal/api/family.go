package api

import (
	"gitlab.com/dentych/dinner-dash/internal/database"
	"gitlab.com/dentych/dinner-dash/internal/models"
)

type FamilyApi struct {
	familyRepo database.FamilyRepo
}

func (api *FamilyApi) Create(user models.User, family models.Family) (models.Family, error) {
	family.Owner = &user
	validationErr := family.Validate()
	if validationErr != nil {
		return models.Family{}, validationErr
	}

	familyID, err := api.familyRepo.Insert(family)
	if err != nil {
		return models.Family{}, err
	}
	family.ID = familyID
	return family, nil
}

func (api *FamilyApi) GetFamily(user models.User, familyID int) (models.Family, error) {
	userInFamily, err := api.familyRepo.UserInFamily(user.ID, familyID)
	if err != nil {
		return models.Family{}, err
	}
	if !userInFamily {
		return models.Family{}, ErrUserNotInFamily
	}
	family, err := api.familyRepo.GetById(familyID)
	if err != nil {
		return models.Family{}, err
	}
	return family, nil
}
