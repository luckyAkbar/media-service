package httpsvc

import (
	"fmt"
	"media-service/internal/config"
	"media-service/internal/usecase"
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

func (s *Service) handleDelete() echo.HandlerFunc {
	type request struct {
		ImageKey  string `json:"image_key"`
		DeleteKey string `json:"delete_key"`
	}

	return func(c echo.Context) error {
		req := request{}
		if err := c.Bind(&req); err != nil {
			logrus.Error(err)
			return ErrInvalidPayload
		}

		if req.ImageKey == "" || req.DeleteKey == "" {
			logrus.Error(ErrInvalidPayload)
			return ErrInvalidPayload
		}

		imageHandler := usecase.NewImageHandler()
		if err := imageHandler.HandleDelete(req.ImageKey, req.DeleteKey); err != nil {
			logrus.Error(err)
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
