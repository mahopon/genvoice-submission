package middleware

import (
	"backend/internal/util"
	"log"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  util.SigningKey,
		TokenLookup: "cookie:access_token",
		ContextKey:  "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(util.Claims)
		},
		SuccessHandler: func(ctx echo.Context) {
			// We do not need to check error as we are sure that the cookie is active and valid
			token := ctx.Get("user").(*jwt.Token)
			if claims, ok := token.Claims.(*util.Claims); !ok {
				log.Println("Claims do not contain required values")
			} else {
				ctx.Set("userid", claims.Subject)
				ctx.Set("role", claims.Role)
			}
		},
	})
}
