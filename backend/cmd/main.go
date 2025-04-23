package main

import (
	"backend/internal/config"
	"backend/internal/middleware"
	repo "backend/internal/repository"
	"backend/internal/route"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	config.LoadConfig()
	repo.InitDB(os.Getenv("DB_URL"))
	e := echo.New()
	route.RegisterRoutes(e)
	// e.Use(middleware.CORS())
	e.Use(middleware.LogResponseTimeMiddleware)
	e.Pre(emw.RemoveTrailingSlash())
	e.Logger.Fatal(e.Start(":8080"))
}
