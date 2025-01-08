package Appmiddleware

import (
	"fmt"

	"github.com/labstack/echo"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("in construction")
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}
