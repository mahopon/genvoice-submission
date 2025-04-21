package middleware

import (
	"backend/internal/util"
	"net/http"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  util.SigningKey,
		TokenLookup: "cookie:access_token",
		ContextKey:  "user",
	})
}

func RoleMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			role := claims["role"].(string)

			// Check if the role is in the allowed roles list
			for _, r := range allowedRoles {
				if role == r {
					return next(c) // Continue to the handler
				}
			}

			// Return forbidden if the role is not allowed
			return echo.NewHTTPError(http.StatusForbidden, "Insufficient privileges")
		}
	}
}
