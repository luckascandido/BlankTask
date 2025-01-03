package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Application struct {
	logger echo.Logger
	server *echo.Echo
}

func main() {

	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}
