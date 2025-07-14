package vecho

import (
	"github.com/labstack/echo/v4"
	"github.com/mymyka/v/docs"
)

type VRoute struct {
	*echo.Route
	R *docs.Route
}

func (r *VRoute) WithDocs(a ...func(*docs.Route)) *VRoute {
	for _, c := range a {
		c(r.R)
	}

	return r
}
