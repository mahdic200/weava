package cmd

import (
	"fmt"
	"os"

	"github.com/mahdic200/weava/Utils/Constants"
	"github.com/spf13/cobra"
)

var showVersion bool

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "weava",
	Long: `
Weava is a Go back-end web framework based on gofiber package for making fast
and powerful restful APIs. This framework is built to give you a great structure for
making fast and powerful web applications in Go language without
any fear.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("Application version %s\n", Constants.VERSION)
			return
		}
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Shows the application version")
}
