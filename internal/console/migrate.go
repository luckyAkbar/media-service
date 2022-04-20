package console

import (
	"image-service/internal/db"
	"image-service/internal/model"

	"github.com/spf13/cobra"
)

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate all database table",
	Long:  "Use this command to initialize your database table for the first time",
	Run:   migrate,
}

func init() {
	RootCMD.AddCommand(migrateCMD)
}

func migrate(cmd *cobra.Command, args []string) {
	db.InitPostgresConn()

	db.DB.AutoMigrate(&model.Image{})
}
