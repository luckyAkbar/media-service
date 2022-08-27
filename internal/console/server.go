package console

import (
	"fmt"
	"media-service/internal/config"
	"media-service/internal/db"
	"media-service/internal/delivery/httpsvc"
	"media-service/internal/repository"
	"media-service/internal/usecase"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCMD = &cobra.Command{
	Use:   "server",
	Short: "start HTTP server",
	Long:  "Start the HTTP server of Media-Service",
	Run:   startServer,
}

func init() {
	RootCMD.AddCommand(serverCMD)
}

func startServer(cmd *cobra.Command, args []string) {
	if err := db.InitPostgresConn(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	sqlDB, err := db.DB.DB()
	if err != nil {
		logrus.Fatal("unable to start server: ", err)
	}

	defer sqlDB.Close()

	videoRepository := repository.NewVideoRepository(db.DB)

	videoUsecase := usecase.NewVideoUsecase(videoRepository)

	server := echo.New()
	server.Pre(middleware.AddTrailingSlash())

	server.Use(middleware.Logger())
	server.Use(middleware.CORS())
	server.Use(middleware.Recover())

	group := server.Group("")

	httpsvc.RouteService(group, videoUsecase)

	server.Start(fmt.Sprintf(":%s", config.ServerPort()))
}
