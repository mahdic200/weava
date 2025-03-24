/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Routes"
	"github.com/spf13/cobra"
)

var port uint16

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the development server (port 8000 by default)",
	Long:  `Serves the dev application`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New()

		if err := Config.Connect(); err != nil {
			fmt.Printf("Could not connect to the database\n")
			fmt.Printf("%v\n", err.Error())
			os.Exit(2)
		}

		Routes.SetupRoutes(app)

		log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().Uint16VarP(&port, "port", "p", 8000, "Sets the port for server")
}
