package cmd

import (
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
	"task_cart/config"
	"task_cart/pkg/db"
)

const path = "migrations"

var migrationCMD = &cobra.Command{
	Use: "migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.MustLoad()
		dbConn := db.MustStartDB(&cfg.DbConf)

		defer db.MustCloseDB(dbConn)

		dbSql, err := dbConn.DB()
		if err != nil {
			panic(err)
		}

		if err = goose.Run(args[0], dbSql, path, args[1:]...); err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCMD.AddCommand(migrationCMD)
}
