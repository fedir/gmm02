package grifts

import (
	"github.com/fedir/gmm02/buffaloweb/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
