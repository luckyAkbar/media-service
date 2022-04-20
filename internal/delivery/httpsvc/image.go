package httpsvc

import (
	"image-service/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) handleSave() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")

		if err != nil {
			return ErrCustomMsgAndStatus(http.StatusBadRequest, err.Error())
		}

		imageHandler := usecase.NewImageHandler(file)
		res, err := imageHandler.HandleUpload()

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}
