package console

import (
	"fmt"
	"image-service/internal/config"
	"image-service/internal/db"
	"image-service/internal/delivery/httpsvc"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCMD = &cobra.Command{
	Use:   "server",
	Short: "start HTTP server",
	Long:  "Start the HTTP server of Image-Service",
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

	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.CORS())
	server.Use(middleware.Recover())

	group := server.Group("")

	httpsvc.RouteService(group)

	server.Start(fmt.Sprintf(":%s", config.ServerPort()))
}
