package main

import (
	"blanktask/common"
	"blanktask/src/api/handlers"

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
	//chama a conexão com o banco
	db, err := common.Mypg()
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
	//Pasaa porta para ativar o endpoints
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.routes(h)

	//porta do servidor
	e.Logger.Fatal(e.Start(":8080"))
}
