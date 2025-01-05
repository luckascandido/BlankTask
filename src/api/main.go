package main

import (
	"blanktask/common"
	"blanktask/src/api/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
)

// cria a extrutura de chamada
type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

// carrega a porta que ira rodar o serviço
type Port struct {
	APP_PORT string `env:"APP_PORT"`
}

func main() {

	e := echo.New()
	//chama a conexão com o banco
	db, err := common.Mypg()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h := handlers.Handler{
		DB: db,
	}
	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}
	//Passaa porta para ativar o endpoints
	fmt.Println(app)
	cfg := Port{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Erro ao carregar variáveis de ambiente:", err)
	}
	appAddress := fmt.Sprintf("localhost:%s", cfg.APP_PORT)
	e.Logger.Fatal(e.Start(appAddress))
}
