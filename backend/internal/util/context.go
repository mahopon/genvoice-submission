package util

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetUUIDFromContext(ctx echo.Context) (parsedUUID uuid.UUID, err error) {
	userId := ctx.Get("userid")
	if parsedUserId, ok := userId.(string); ok {
		parsedUUID, err = uuid.Parse(parsedUserId)
	} else {
		err = errors.New("invalid user id")
	}

	return parsedUUID, err
}
