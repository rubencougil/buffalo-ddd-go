package actions

import (
	"encoding/json"
	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
	"github.com/rubencougil/coke/models"
	"github.com/rubencougil/coke/src/users"
	"io/ioutil"
	"net/http"
)

func ApiUserHandler(c buffalo.Context) error {

	usersApp := &users.App{Repository: &users.PostgresRepository{DB:models.DB}}

	c.Logger().Info()

	switch c.Request().Method {
	case http.MethodGet:
		return get(c, usersApp)
	case http.MethodPost:
		return create(c, usersApp)
	}
	return nil
}

func create(c buffalo.Context, usersApp users.Application) error {
	c.Logger().Info(c.Request().Body)
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.Render(500, r.Plain(err.Error()))
	}
	var user users.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return c.Render(500, r.Plain(err.Error()))
	}
	err = usersApp.Create(user)
	if err != nil {
		return c.Render(500, r.Plain(err.Error()))
	}
	return c.Render(202, r.JSON(user))
}

func get(c buffalo.Context, usersApp users.Application) error {
	id, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		return c.Render(500, r.Plain(err.Error()))
	}
	user, err := usersApp.Get(id)
	if err != nil {
		return c.Render(500, r.Plain(err.Error()))
	}
	return c.Render(200, r.JSON(user))
}
