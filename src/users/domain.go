package users

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type Created struct {
	User User
}

type Repository interface {
	Create(user User) (err error)
	Get(ID uuid.UUID) (user *User, err error)
	GetAll() (users []*User, err error)
}

type Application interface {
	Create(user User) (err error)
	Get(ID uuid.UUID) (user *User, err error)
	GetAll() (users []*User, err error)
}
