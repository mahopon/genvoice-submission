package controller

import (
	model "backend/internal/model/user"
	"backend/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{s}
}

func (c *UserController) LoginUser(ctx echo.Context) error {
	var user model.LoginUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}

	resp := c.UserService.LoginUser(user)
	if resp.StatusCode != 200 {
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) RegisterUser(ctx echo.Context) error {
	var user model.CreateUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}
	resp := c.UserService.RegisterUser(user)
	if resp.StatusCode != 200 {
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")
	resp := c.UserService.DeleteUser(id)
	if resp.StatusCode != 200 {
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	var user model.UpdateUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}
	resp := c.UserService.UpdateUser(id, user)
	if resp.StatusCode != 200 {
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	return ctx.JSON(http.StatusOK, resp)
}
