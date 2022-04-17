package httpsvc

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) handleSave() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.RealIP())
		fmt.Println(c.Request().Header.Values("X-Real-IP"))
		return c.JSON(http.StatusOK, "hello world")
	}
}
