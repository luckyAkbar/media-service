package httpsvc

import (
	"media-service/internal/model"

	"github.com/labstack/echo/v4"
)

type Service struct {
	group        *echo.Group
	videoUsecase model.VideoUsecase

	// TODO use DI for image service
}

func RouteService(group *echo.Group, videoUsecase model.VideoUsecase) {
	srv := &Service{
		group:        group,
		videoUsecase: videoUsecase,
	}

	srv.initRoutes()
	srv.initVideoServiceRoutes()
}

func (s *Service) initRoutes() {
	s.group.GET("/:imageKey/", s.handleGet())
	s.group.POST("/", s.handleSave())
	s.group.PATCH("/", s.handleUpdate())
	s.group.DELETE("/", s.handleDelete())
}

func (s *Service) initVideoServiceRoutes() {
	s.group.POST("/video/", s.handleUploadVideo())
	s.group.GET("/video/:id/", s.handleDownloadVideo())
}
