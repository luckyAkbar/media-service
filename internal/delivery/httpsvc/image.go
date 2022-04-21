package httpsvc

import (
	"fmt"
	"image-service/internal/config"
	"image-service/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (s *Service) handleSave() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")

		if err != nil {
			logrus.Error(err)
			return ErrCustomMsgAndStatus(http.StatusBadRequest, err.Error())
		}

		imageHandler := usecase.NewImageHandler()
		res, err := imageHandler.HandleUpload(file)

		if err != nil {
			logrus.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (s *Service) handleGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		imageKey := c.Param("imageKey")

		imageHandler := usecase.NewImageHandler()
		data, err := imageHandler.HandleGet(imageKey)
		if err != nil {
			logrus.Error(err)
			return ErrImageNotFound
		}

		return c.File(fmt.Sprintf("%s%s", config.ImageStoragePath(), data.Name))
	}
}

func (s *Service) handleUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			logrus.Error(err)
			return ErrCustomMsgAndStatus(http.StatusBadRequest, err.Error())
		}

		imageKey := c.FormValue("imageKey")
		updateKey := c.FormValue("updateKey")

		if imageKey == "" || updateKey == "" {
			return ErrInvalidPayload
		}

		imageHandler := usecase.NewImageHandler()
		keys, err := imageHandler.HandleUpdate(file, imageKey, updateKey)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, keys)
	}
}
