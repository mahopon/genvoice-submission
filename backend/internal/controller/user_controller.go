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

	resp, err := c.UserService.LoginUser(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, resp)
}
