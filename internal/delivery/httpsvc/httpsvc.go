package httpsvc

import "github.com/labstack/echo/v4"

type Service struct {
	group *echo.Group
}

func RouteService(group *echo.Group) {
	srv := &Service{
		group: group,
	}

	srv.initRoutes()
}

func (s *Service) initRoutes() {
	s.group.GET("/:imageKey", s.handleGet())
	s.group.POST("/", s.handleSave())
	s.group.PATCH("/", s.handleUpdate())
	s.group.DELETE("/", s.handleDelete())
}
