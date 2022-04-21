package httpsvc

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrImageNotFound  = echo.NewHTTPError(http.StatusNotFound, "image not found")
	ErrInvalidPayload = echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
)

func ErrCustomMsgAndStatus(status int, msg string) *echo.HTTPError {
	return echo.NewHTTPError(status, msg)
}
