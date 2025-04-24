package middleware

import (
	"backend/internal/util"
	"log"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  util.SigningKey,
		TokenLookup: "cookie:access_token",
		ContextKey:  "user",
		SuccessHandler: func(ctx echo.Context) {
			// We do not need to check error as we are sure that the cookie is active and valid
			userClaims := ctx.Get("user")
			if claims, ok := userClaims.(*util.Claims); !ok {
				log.Println("This ain't right!")
			} else {
				ctx.Set("userid", claims.Subject)
				ctx.Set("role", claims.Role)
			}
		},
	})
}
