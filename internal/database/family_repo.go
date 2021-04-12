package database

//go:generate mockgen -destination mocks/family_repo.go . FamilyRepo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/dentych/dinner-dash/internal/models"
	"strings"
)

type FamilyRepo interface {
	GetById(ctx context.Context, id int) (Family, error)
	GetByInvitationID(ctx context.Context, invitationID string) (Family, error)
	Insert(ctx context.Context, familyName string, userID string) (int, error)
	Update(ctx context.Context, family Family, fieldsToUpdate ...string) error
	Delete(ctx context.Context, family models.Family) error
	UserInFamily(ctx context.Context, userID string, familyID int) (bool, error)
	AddMember(ctx context.Context, familyID int, userID string) error
}

type Family struct {
	ID           int
	Name         *string
	InvitationID *string `db:"invitation_id"`
}

type familyRepo struct {
	db *sqlx.DB
}

func NewFamilyRepo(db *sqlx.DB) FamilyRepo {
	return &familyRepo{db: db}
}

func (f *familyRepo) GetById(ctx context.Context, id int) (Family, error) {
	var family Family
	err := f.db.Get(&family, "select * from public.family where id = $1", id)
	return family, err
}

func (f *familyRepo) Insert(ctx context.Context, familyName string, userID string) (int, error) {
	tx, err := f.db.Beginx()
	if err != nil {
		return 0, err
	}

	var familyId int
	err = tx.Get(&familyId, "insert into public.family (name) values ($1) returning id", familyName)
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec("update public.user set family_id = $1 where id = $2", familyId, userID)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, nil
	}
	return familyId, nil
}

func (f *familyRepo) Update(ctx context.Context, family Family, fieldsToUpdate ...string) error {
	var sqlKeys []string
	var sqlValues []interface{}
	for _, v := range fieldsToUpdate {
		key := strings.ToLower(v)
		switch key {
		case "name":
			sqlKeys = append(sqlKeys, "name")
			sqlValues = append(sqlValues, family.Name)
		case "invitationid":
			sqlKeys = append(sqlKeys, "invitation_id")
			sqlValues = append(sqlValues, family.InvitationID)
		}
	}

	if len(sqlKeys) == 0 {
		return nil
	}

	sql := "update public.family set"
	for k, v := range sqlKeys {
		sql += fmt.Sprintf(" %s=$%d,", v, k+1)
	}
	sql = sql[:len(sql)-1]
	sql += fmt.Sprintf(" where id = $%d", len(sqlKeys)+1)
	sqlValues = append(sqlValues, family.ID)

	_, err := f.db.Exec(sql, sqlValues...)
	if err != nil {
		return err
	}
	return nil
}

func (f *familyRepo) Delete(ctx context.Context, family models.Family) error {
	panic("implement me")
}

func (f *familyRepo) UserInFamily(ctx context.Context, userID string, familyID int) (bool, error) {
	var familyCount int
	err := f.db.Get(&familyCount, "select count(*) from public.user where id = $1 and family_id = $2", userID, familyID)
	return familyCount > 0, err
}

func (f *familyRepo) GetByInvitationID(ctx context.Context, invitationID string) (Family, error) {
	var family Family
	err := f.db.Get(&family, "select * from public.family where invitation_id = $1", invitationID)
	if err != nil {
		return Family{}, err
	}

	return family, nil
}

func (f *familyRepo) AddMember(ctx context.Context, familyID int, userID string) error {
	_, err := f.db.Exec("update public.user set family_id = $1 where id = $2", familyID, userID)
	if err != nil {
		return err
	}
	return nil
}
