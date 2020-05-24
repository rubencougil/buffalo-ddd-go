package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rubencougil/coke/models"
	"github.com/rubencougil/coke/src/users"
)

func ApiUsersHandler(c buffalo.Context) error {

	usersApp := &users.App{Repository: &users.PostgresRepository{DB:models.DB}}

	users, err := usersApp.GetAll()
	if err != nil {
		c.Logger().Error(err)
		return c.Render(200, r.Plain(err.Error()))
	}
	return c.Render(200, r.JSON(users))
}
