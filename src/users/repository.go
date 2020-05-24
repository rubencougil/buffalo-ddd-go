package users

import (
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/rubencougil/coke/models"
	"time"
)

type PostgresRepository struct {
	DB *pop.Connection
}

func (r PostgresRepository) Create(user User) error {
	if err := r.DB.Save(models.User{
		ID:        user.ID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}); err != nil {
		return errors.Wrap(err, "Error saving User into the DB")
	}

	return nil
}

func (r PostgresRepository) Get(ID uuid.UUID) (user *User, err error) {
	userModel := models.User{}
	if err = r.DB.Find(&userModel, ID); err != nil {
		return nil, errors.Wrap(err, "Cannot find User in Postgres DB")
	}
	return &User{
		ID:   userModel.ID,
		Name: userModel.Name,
	}, nil
}

func (r PostgresRepository) GetAll() (users []*User, err error) {
	usersModel := &models.Users{}
	if err = r.DB.All(usersModel); err != nil {
		return nil, errors.Wrap(err, "Cannot get all Users in Postgres DB")
	}
	for _, userModel := range *usersModel {
		users = append(users, &User{
			ID:   userModel.ID,
			Name: userModel.Name,
		})
	}
	return users, nil
}


