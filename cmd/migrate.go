package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"url-shortener/config"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migration of database",
		Long: `Migration of mysql database!

Flags:
	--action set action
	--config set config address (default is ./config)
	--folder set migration folder (default is ./migrations)
	`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Init(configPath, "yaml")
			if err != nil {
				panic(err)
			}

			dbCfg := cfg.DB

			sourceAddress := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
				dbCfg.Username,
				dbCfg.Password,
				dbCfg.Host,
				dbCfg.Port,
			)

			tableStarter, err := sql.Open(dbCfg.Driver, sourceAddress)
			if err != nil {
				panic(err)
			}

			if _, err = tableStarter.Exec("CREATE DATABASE IF NOT EXISTS " + dbCfg.DB); err != nil {
				panic(err)
			}

			dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
				dbCfg.Username,
				dbCfg.Password,
				dbCfg.Host,
				dbCfg.Port,
				dbCfg.DB)
			db, err := sql.Open(dbCfg.Driver, dataSourceName)
			if err != nil {
				panic(err)
			}

			driver, err := mysql.WithInstance(db, &mysql.Config{})
			if err != nil {
				panic(err)
			}

			m, err := migrate.NewWithDatabaseInstance(
				fmt.Sprintf("file://%s", migrationFolder),
				dbCfg.Driver,
				driver,
			)
			if err != nil {
				panic(err)
			}

			if action == "up" {
				err = m.Up()
			} else {
				err = m.Down()
			}

			if err != nil {
				panic(err)
			}
		},
	}
	action          string
	configPath      string
	migrationFolder string
)

func init() {
	migrateCmd.Flags().StringVar(&configPath, "config", "./config", "set config path")
	migrateCmd.Flags().StringVar(&action, "action", "", "set action (either up or down)")
	migrateCmd.Flags().StringVar(&migrationFolder, "folder", "./migrations", "set migration files folder")
	rootCmd.AddCommand(migrateCmd)
}
