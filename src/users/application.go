package users

import (
	"fmt"
	"github.com/gobuffalo/events"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type App struct {
	Repository Repository
}

func init() {
	_, err := events.Listen(func(e events.Event) {
		fmt.Println(e.Kind)
		if e.Kind == "coke:users:create" {
			fmt.Println(e.Payload)
		}
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (c App) Create(user User) (err error) {
	err = c.Repository.Create(user)

	if err != nil {
		return errors.Wrap(err, "Error creating User")
	}

	err = events.EmitPayload("coke:users:create", user)
	if err != nil {
		return errors.Wrap(err, "Error publishing Create User event")
	}

	return
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
