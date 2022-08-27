package httpsvc

import (
	"fmt"
	"media-service/internal/config"
	"media-service/internal/usecase"
	"net/http"
	"strconv"

	"github.com/kumparan/go-utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (s *Service) handleUploadVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		file, err := c.FormFile("video")
		if err != nil {
			logrus.Warn("failed to read vide from form file: ", err)
			return ErrBadRequest
		}

		video, err := s.videoUsecase.Upload(c.Request().Context(), file, title)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":   utils.DumpIncomingContext(c.Request().Context()),
				"title": title,
				"file":  utils.Dump(file),
			}).Error(err)

			return ErrInternal
		case usecase.ErrMimeTypeNotFound:
			return ErrMimeTypeNotFound
		case usecase.ErrMimeTypeIsForbidden:
			return ErrMimeTypeIsForbidden
		case usecase.ErrFileSizeToLarge:
			return ErrFileTooLarge
		case nil:
			return c.JSON(http.StatusOK, video)
		}
	}
}

func (s *Service) handleDownloadVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		if param == "" {
			return ErrBadRequest
		}

		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return ErrBadRequest
		}

		video, err := s.videoUsecase.Download(c.Request().Context(), id)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx": utils.DumpIncomingContext(c.Request().Context()),
				"id":  id,
			}).Error(err)

			return ErrInternal
		case usecase.ErrNotFound:
			return ErrNotFound
		case nil:
			c.Response().Header().Set("Content-Type", video.MimeType)
			return c.File(fmt.Sprintf("%s%s", config.VideoStoragePath(), video.Filename))
		}
	}
}
