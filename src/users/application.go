package users

import (
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type App struct {
	Repository Repository
}

func (c App) Create(user User) error {
	return c.Repository.Create(user)
}

func (c App) Get(ID uuid.UUID) (user *User, err error) {
	user, err = c.Repository.Get(ID)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot get User")
	}
	return user, nil
}

func (c App) GetAll() (user []*User, err error) {
	return c.Repository.GetAll()
}
