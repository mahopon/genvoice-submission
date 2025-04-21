package main

import (
	"backend/internal/config"
	repo "backend/internal/repository"
	"backend/internal/route"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	config.LoadConfig()
	repo.InitDB(os.Getenv("DB_URL"))
	e := echo.New()
	route.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
