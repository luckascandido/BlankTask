package main

import (
	"blanktask/internal/database"
	"blanktask/src/api/handlers"
	Appmiddleware "blanktask/src/api/middleware"
	"context"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// cria a estrutura de chamada.
type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {

	e := echo.New()
	ctx := context.Background()
	//chama a conexão com o banco
	db, err := database.Mypg(ctx)
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	h := handlers.Handler{
		DB: db,
	}
	// faz a configuração geral de conexão
	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}
	// Carrega as migraçôes
	if err := database.Migration_up(ctx); err != nil {
		e.Logger.Fatal("Erro durante a migração:", err)
	}
	//Pasaa porta para ativar o endpoints
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(Appmiddleware.CustomMiddleware)
	app.routes(h)

	//porta do servidor
	e.Logger.Fatal(e.Start(":8080"))
}
