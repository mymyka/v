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
	echoRoute := g.Group.GET(path, h, m...)

	docsRoute := docs.Route{
		Method: "get",
		Path:   path,
	}

	vRoute := VRoute{
		Route: echoRoute,
		R:     &docsRoute,
	}

	return &vRoute
}
