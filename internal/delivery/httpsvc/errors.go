package httpsvc

import "github.com/labstack/echo/v4"

func ErrCustomMsgAndStatus(status int, msg string) *echo.HTTPError {
	return echo.NewHTTPError(status, msg)
}
