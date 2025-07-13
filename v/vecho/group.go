package vecho

import (
	"github.com/labstack/echo/v4"
	"github.com/mymyka/v/v/docs"
)

type VGroup struct {
	*echo.Group
	DocsGroup *docs.Group
}

func (g *VGroup) GET(path string, h func(echo.Context) error, m ...echo.MiddlewareFunc) *VRoute {
	r := g.Group.GET(path, h, m...)
	return NewVRoute(r)
}
