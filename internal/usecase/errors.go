package usecase

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrServerFailedToSaveImage = echo.NewHTTPError(http.StatusInternalServerError, "server failed to save the uploaded image")
	ErrMimeTypeNotFound        = echo.NewHTTPError(http.StatusBadRequest, "Mimetype is not found")
)

func ErrMimeTypeForbidden(mimetype string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("mimetype: %s is forbidden", mimetype))
}

func ErrFileToLarge(actual, max int64) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("Image too large with size: %d, max only: %d", actual, max))
}
