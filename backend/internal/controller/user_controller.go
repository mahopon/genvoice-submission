package controller

import (
	model "backend/internal/model/user"
	"backend/internal/service"
	"backend/internal/util"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{s}
}

func (c *UserController) GetUser(ctx echo.Context) error {
	id := ctx.Param("id")
	resp := c.UserService.GetUser(id)
	if resp == nil {
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) LoginUser(ctx echo.Context) error {
	var user model.LoginUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}

	dbUser, err := c.UserService.LoginUser(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	accessToken, _ := util.CreateAccessToken(dbUser.ID, dbUser.Role)
	refreshToken, _ := util.CreateRefreshToken(dbUser.ID, dbUser.Role)

	// Set refresh token in HttpOnly cookie
	ctx.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true, // Set true in production (HTTPS)
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	// Set access token in HttpOnly cookie
	ctx.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Secure:   true, // Set true in production (HTTPS)
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(15 * time.Minute),
	})

	return ctx.JSON(http.StatusOK, echo.Map{
		"id":           dbUser.ID,
		"name":         dbUser.Name,
		"access_token": accessToken,
	})
}

func (c *UserController) Refresh(ctx echo.Context) error {
	// Get the refresh token from the cookie
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		// No refresh token found, user is not authenticated
		return echo.ErrUnauthorized
	}

	// Parse the refresh token
	token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
		return []byte(util.SigningKey), nil
	})
	if err != nil || !token.Valid {
		// If token is invalid, log out the user by clearing both tokens
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	// Extract claims from the refresh token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["typ"] != "refresh" {
		// If it's not a refresh token, log out the user
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	// Check if the refresh token is expired
	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		// If the refresh token is expired, log out the user
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	// Extract user info from claims
	userIDStr, ok := claims["sub"].(string)
	role, ok2 := claims["role"].(string)
	if !ok || !ok2 {
		// If user info is missing, log out the user
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	// Parse the user ID to UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		// Invalid user ID, log out the user
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	// Create a new access token (short-lived)
	accessToken, _ := util.CreateAccessToken(userID, role)

	// Set the new access token cookie (short-lived)
	ctx.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Secure:   false, // Set true in production (HTTPS)
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	// Return the new access token in the response
	return ctx.JSON(http.StatusOK, echo.Map{
		"access_token": accessToken,
	})
}

func (c *UserController) Logout(ctx echo.Context) error {
	util.ClearTokens(ctx)
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Logged out"})
}

func (c *UserController) RegisterUser(ctx echo.Context) error {
	var user model.CreateUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}
	err := c.UserService.RegisterUser(user)
	if err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *UserController) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.UserService.DeleteUser(id)
	if err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *UserController) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	var user model.UpdateUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}
	err := c.UserService.UpdateUser(id, user)
	if err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *UserController) CheckAuthStatus(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
