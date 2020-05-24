package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rubencougil/coke/models"
	"github.com/rubencougil/coke/src/users"
	"net/http"
)

func UsersHandler(c buffalo.Context) error {
	usersApp := &users.App{Repository: &users.PostgresRepository{DB:models.DB}}
	users, err := usersApp.GetAll()
	if err != nil {
		c.Logger().Error(err)
		return c.Render(http.StatusNotFound, r.HTML("index.html"))
	}
	c.Set("users", users)
	return c.Render(http.StatusNotFound, r.HTML("users.html"))
}
