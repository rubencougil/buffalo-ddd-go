package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rubencougil/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
