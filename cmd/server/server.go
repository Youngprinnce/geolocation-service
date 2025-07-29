package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/youngprinnce/geolocation-service/config"
	"github.com/youngprinnce/geolocation-service/internal/logger"
	"github.com/youngprinnce/geolocation-service/internal/postgres"
)

func StartServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start the API server",
		Long:  `Start the HTTP API server`,
		Run: func(cmd *cobra.Command, args []string) {
			configFile, _ := cmd.Flags().GetString("config")
			conf := config.LoadConfig(configFile)

			logger.Initialize()

			if err := postgres.Load(conf); err != nil {
				logger.Fatal(fmt.Sprintf("Failed to initialize postgres: %v", err))
			}

			router := RegisterRoutes(conf)

			log.WithField("port", conf.Server.Listen).Info("Starting server")
			if err := router.Run(conf.Server.Listen); err != nil {
				logger.Fatal(fmt.Sprintf("Failed to start server: %v", err))
			}
		},
	}
}
