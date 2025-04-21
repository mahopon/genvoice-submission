package main

import (
	"backend/internal/config"
	"backend/internal/middleware"
	repo "backend/internal/repository"
	"backend/internal/route"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	repo.InitDB(os.Getenv("DB_URL"))
	e := echo.New()
	route.RegisterRoutes(e)
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8080"))
}
