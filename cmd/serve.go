package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/http/handler"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start application",
		Long: `Start application with this command.
Do not forget to migrate before start.

Flag:
	--config set config path (default is ./config)`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Init(configPath, "yaml")
			if err != nil {
				panic(err)
			}

			db, err := database.Init(cfg.DB)
			if err != nil {
				panic(err)
			}

			e := echo.New()

			sbmt := handler.Submit{DB: db}
			e.POST("/submit", sbmt.SubmitHandler)

			rcv := handler.Receive{DB: db}
			//e.HEAD("/receive/", rcv.ReceiveHandler)
			e.Any("/receive/:", rcv.ReceiveHandler)

			err = e.Start(cfg.Server.Host + ":" + cfg.Server.Port)
			if err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVar(&configPath, "config", "./config", "set config path (default is ./config)")
}
