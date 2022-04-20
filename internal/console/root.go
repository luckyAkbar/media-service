package console

import (
	"image-service/internal/config"
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	Use: "image-service",
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	setupLogger()
}

func setupLogger() {
	formatter := runtime.Formatter{
		ChildFormatter: &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Line:    true,
		File:    true,
		Package: true,
	}

	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(config.LogLevel())

	if err != nil {
		logLevel = logrus.DebugLevel
	}

	logrus.SetLevel(logLevel)
}
