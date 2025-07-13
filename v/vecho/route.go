package vecho

import "github.com/labstack/echo/v4"

type VRoute struct {
	*echo.Route
}

func NewVRoute(r *echo.Route) *VRoute {
	return &VRoute{
		Route: r,
	}
}

func (r *VRoute) WithDocs(a ...any) {}
