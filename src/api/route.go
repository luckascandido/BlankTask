package main

import (
	"blanktask/src/api/handlers"
)

func (app *Application) routes(h handlers.Handler) {
	app.server.GET("/", h.HealthCheck)
	app.server.POST("/register", h.RegisterHandler)
}
