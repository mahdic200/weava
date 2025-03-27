/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/mahdic200/weava/Config"
	"github.com/spf13/cobra"
)

var port uint16

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the development server (port 8000 by default)",
	Long:  `Serves the dev application`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(Config.App.Listen(fmt.Sprintf(":%v", port)))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().Uint16VarP(&port, "port", "p", 8000, "Sets the port for server")
}
