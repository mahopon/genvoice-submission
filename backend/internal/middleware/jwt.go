package middleware

import (
	"backend/internal/util"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
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
			token, err := c.Cookie("access_token")
			if err != nil {
				log.Println("No token?!1")
				return echo.ErrUnauthorized
			}

			parsedToken, err := util.ValidateJWT(token.Value)
			if err != nil {
				log.Printf("Invalid token?!2 %v,", err)
				return echo.ErrUnauthorized
			}

			role := parsedToken.Role

			// Check if the role is in the allowed roles list
			for _, r := range allowedRoles {
				if strings.EqualFold(role, r) {
					return next(c) // Continue to the handler
				}
			}

			// Return forbidden if the role is not allowed
			return echo.NewHTTPError(http.StatusForbidden, "Insufficient privileges")
		}
	}
}
