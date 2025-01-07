package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HealthCheck struct {
	Health bool `json:"health"`
}

// função teste.
func (h *Handler) HealthCheck(c echo.Context) error {
	healthCheckStruct := HealthCheck{
		Health: true,
	}

	return c.JSON(http.StatusOK, healthCheckStruct)

}
