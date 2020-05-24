package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
)

func ApiUserHandler(c buffalo.Context) error {
	c.Logger().Info(fmt.Sprintf("%s: %s", c.Request().Method, c.Param("id")))
	return nil
}