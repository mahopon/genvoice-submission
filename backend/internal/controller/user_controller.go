package controller

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/internal/util"
	"log"
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
	user, err := c.UserService.GetUser(id)
	if err != nil || user == nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}
	return ctx.JSON(http.StatusOK, user)
}

func (c *UserController) LoginUser(ctx echo.Context) error {
	// Check if a valid JWT is present in the cookies
	tokenCookie, err := ctx.Request().Cookie("access_token")
	if err == nil {
		// Validate the JWT token
		claims, err := util.ValidateJWT(tokenCookie.Value)
		if err == nil {
			// JWT is valid, skip the login and return the user details
			dbUser, err := c.UserService.GetUser(claims.Subject)
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
			}

			// Return the user data
			return ctx.JSON(http.StatusOK, echo.Map{
				"id":   dbUser.ID,
				"name": dbUser.Name,
			})
		}
	}

	var user model.LoginUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	dbUser, err := c.UserService.LoginUser(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	accessToken, _ := util.CreateAccessToken(dbUser.ID, dbUser.Role)
	refreshToken, _ := util.CreateRefreshToken(dbUser.ID, dbUser.Role)

	c.setTokenCookies(ctx, accessToken, refreshToken)

	return ctx.JSON(http.StatusOK, echo.Map{
		"id":   dbUser.ID,
		"name": dbUser.Name,
	})
}

func (c *UserController) Refresh(ctx echo.Context) error {
	cookie, err := ctx.Cookie("refresh_token")
	var tokenToRefresh bool = false // False for access, true for refresh
	if err != nil {
		cookie, err = ctx.Cookie("access_token")
		if err == nil {
			tokenToRefresh = true
		} else {
			return echo.ErrUnauthorized
		}
	}

	token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
		return []byte(util.SigningKey), nil
	})
	if err != nil || !token.Valid {
		log.Printf("Err: %v", err)
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		log.Printf("Err: %v, exp", err)
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	userIDStr, ok := claims["sub"].(string)
	role, ok2 := claims["role"].(string)
	if !ok || !ok2 {
		log.Printf("Err: %v, sub or role", err)
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		log.Printf("Err: %v, uuid", err)
		util.ClearTokens(ctx)
		return echo.ErrUnauthorized
	}

	if tokenToRefresh {
		accessToken, _ := util.CreateRefreshToken(userID, role)
		ctx.SetCookie(&http.Cookie{
			Name:     "refresh_token",
			Value:    accessToken,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
		})
	} else {
		accessToken, _ := util.CreateAccessToken(userID, role)
		ctx.SetCookie(&http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
			Expires:  time.Now().Add(15 * time.Minute),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"role": role,
	})
}

func (c *UserController) Logout(ctx echo.Context) error {
	util.ClearTokens(ctx)
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Logged out"})
}

func (c *UserController) RegisterUser(ctx echo.Context) error {
	var user model.CreateUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}
	if err := c.UserService.RegisterUser(user); err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *UserController) DeleteUser(ctx echo.Context) error {
	if err := c.UserService.DeleteUser(ctx.Param("id")); err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *UserController) UpdateUser(ctx echo.Context) error {
	var user model.UpdateUserRequest
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}
	if err := c.UserService.UpdateUser(ctx.Param("id"), user); err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *UserController) CheckAuthStatus(ctx echo.Context) error {
	tokenCookie, err := ctx.Request().Cookie("access_token")
	if err == nil {
		claims, err := util.ValidateJWT(tokenCookie.Value)
		if err == nil {
			_, err := ctx.Request().Cookie("refresh_token")
			if err != nil {
				return ctx.Redirect(http.StatusFound, "/api/user/refresh")
			}

			return ctx.JSON(http.StatusOK, echo.Map{"role": claims.Role, "user_id": claims.Subject})
		}
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid access token"})
	}

	_, err = ctx.Request().Cookie("refresh_token")
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Not authenticated"})
	}

	return ctx.Redirect(http.StatusFound, "/api/user/refresh")
}

func (c *UserController) setTokenCookies(ctx echo.Context, accessToken, refreshToken string) {
	ctx.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})
	ctx.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(15 * time.Minute),
	})
}
