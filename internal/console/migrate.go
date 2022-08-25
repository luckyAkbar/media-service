package console

import (
	"media-service/internal/config"
	"media-service/internal/db"
	"media-service/internal/model"
	"strconv"

	"github.com/kumparan/go-utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	migrate "github.com/rubenv/sql-migrate"
)

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate all database table",
	Long:  "Use this command to initialize your database table for the first time",
	Run:   runMigrate,
}

func init() {
	migrateCMD.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrateCMD.PersistentFlags().String("direction", "up", "migration direction")
	RootCMD.AddCommand(migrateCMD)
}

func runMigrate(cmd *cobra.Command, args []string) {
	if err := db.InitPostgresConn(); err != nil {
		logrus.Fatal(err)
	}

	if err := db.DB.AutoMigrate(&model.Image{}); err != nil {
		logrus.Fatal(err)
	}

	direction := cmd.Flag("direction").Value.String()
	stepStr := cmd.Flag("step").Value.String()
	step, err := strconv.Atoi(stepStr)
	if err != nil {
		logrus.WithField("stepStr", stepStr).Fatal("Failed to parse step to int: ", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./db/migration",
	}

	migrate.SetTable("schema_migrations")

	pgdb, err := db.DB.DB()
	if err != nil {
		logrus.WithField("DatabaseDSN", config.PostgresDSN()).Fatal("failed to run migration")
	}

	var n int
	if direction == "down" {
		n, err = migrate.ExecMax(pgdb, "postgres", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(pgdb, "postgres", migrations, migrate.Up, step)
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"migrations": utils.Dump(migrations),
			"direction":  direction}).
			Fatal("Failed to migrate database: ", err)
	}

	logrus.Infof("Applied %d migrations!\n", n)
}
