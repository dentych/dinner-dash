package database

//go:generate mockgen -destination mocks/family_repo.go . FamilyRepo

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/dentych/dinner-dash/internal/models"
)

type FamilyRepo interface {
	GetById(id int) (models.Family, error)
	Insert(family models.Family) (int, error)
	Update(family models.Family) error
	Delete(family models.Family) error
	UserInFamily(userID int, familyID int) (bool, error)
}

type familyRepo struct {
	db *sqlx.DB
}

func (f familyRepo) GetById(id int) (models.Family, error) {
	//db.Get()
	panic("implement me")
}

func (f familyRepo) Insert(family models.Family) (int, error) {
	panic("implement me")
}

func (f familyRepo) Update(family models.Family) error {
	panic("implement me")
}

func (f familyRepo) Delete(family models.Family) error {
	panic("implement me")
}

func (f familyRepo) UserInFamily(userID int, familyID int) (bool, error) {
	panic("implement me")
}
