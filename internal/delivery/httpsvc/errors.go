package httpsvc

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrImageNotFound       = echo.NewHTTPError(http.StatusNotFound, "image not found")
	ErrInvalidPayload      = echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
	ErrBadRequest          = echo.NewHTTPError(http.StatusBadRequest, "bad request")
	ErrInternal            = echo.NewHTTPError(http.StatusInternalServerError, "internal error")
	ErrMimeTypeNotFound    = echo.NewHTTPError(http.StatusBadRequest, "mime type not found")
	ErrMimeTypeIsForbidden = echo.NewHTTPError(http.StatusForbidden, "mime type forbidden")
	ErrFileTooLarge        = echo.NewHTTPError(http.StatusBadRequest, "file too large")
)

func ErrCustomMsgAndStatus(status int, msg string) *echo.HTTPError {
	return echo.NewHTTPError(status, msg)
}
