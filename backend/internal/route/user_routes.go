package route

import (
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/service"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(g *echo.Group) {
	userService := service.NewUserService()
	userController := controller.NewUserController(userService)

	// Public routes
	// POST /api/user/login
	g.POST("/login", userController.LoginUser)
	// POST /api/user/register
	g.POST("/register", userController.RegisterUser)

	// Protected routes
	// GET /api/user/:id
	g.GET("/:id", userController.GetUser, middleware.JWTMiddleware())
	// GET /api/user/auth
	g.GET("/auth", userController.CheckAuthStatus, middleware.JWTMiddleware())
	// PATCH /api/user/update/:id
	g.PATCH("/update/:id", userController.UpdateUser, middleware.JWTMiddleware())
	// POST /api/user/logout
	g.POST("/logout", userController.Logout, middleware.JWTMiddleware())

	// Admin-Only routes
	// DELETE /api/user/delete/:id
	g.DELETE("/delete/:id", userController.DeleteUser, middleware.JWTMiddleware(), middleware.RoleMiddleware("ADMIN"))
}
