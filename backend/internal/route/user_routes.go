package route

import (
	"backend/internal/controller"
	"backend/internal/service"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(g *echo.Group) {
	userService := service.NewUserService()
	userController := controller.NewUserController(userService)

	g.POST("/login", userController.LoginUser)
}
