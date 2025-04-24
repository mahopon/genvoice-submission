package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func LogResponseTimeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start)
		log.Printf("Request %s %s took %v", c.Request().Method, c.Request().URL.Path, duration)
		return err
	}
}
