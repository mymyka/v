package vecho

import (
	"github.com/labstack/echo/v4"
	"github.com/mymyka/v/v/docs"
)

type VEcho struct {
	*echo.Echo
	Docs docs.RootDocs
}

func New(e *echo.Echo) *VEcho {
	return &VEcho{
		Echo: e,
		Docs: docs.RootDocs{},
	}
}

func (e *VEcho) Group(prefix string, m ...echo.MiddlewareFunc) (g *VGroup) {
	vg := docs.Group{
		Path: prefix,
	}

	e.Docs.Groups = append(e.Docs.Groups, &vg)

	return &VGroup{
		Group:     e.Echo.Group(prefix, m...),
		DocsGroup: &vg,
	}
}
