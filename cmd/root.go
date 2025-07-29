package cmd

import (
	"github.com/spf13/cobra"
	"github.com/youngprinnce/geolocation-service/cmd/server"
)

var rootCmd = &cobra.Command{
	Use:   "geolocation-service",
	Short: "Geolocation Service API",
	Long:  `A RESTful API for managing geolocated stations and finding nearest locations`,
}

func Execute() {
	rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "config filename")
	rootCmd.AddCommand(server.StartServerCmd())
	cobra.CheckErr(rootCmd.Execute())
}
