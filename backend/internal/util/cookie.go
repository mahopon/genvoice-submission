package util

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func ClearTokens(ctx echo.Context) {
	// Clear refresh token
	ctx.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Expired in the past
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// Clear access token
	ctx.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Expired in the past
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
}
