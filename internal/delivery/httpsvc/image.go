package httpsvc

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) handleSave() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	}
}
