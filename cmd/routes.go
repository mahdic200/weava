package cmd

import (
	"fmt"
	"strings"

	"github.com/mahdic200/weava/Config"
	"github.com/spf13/cobra"
)

type Route struct {
	Name   string
	Method string
	Path   string
}

// routesCmd represents the routes command
var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Shows defined routes list",
	Long:  `shows a list of defined routes in your application .`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, route := range Config.App.GetRoutes() {
			if route.Path != "/*" && route.Name != "" {
				if strings.Contains(route.Name, "index") && route.Method == "HEAD" {
					continue
				}
				// fmt.Printf("path : %#v, name : %#v, method : %#v\n", route.Path, route.Name, route.Method)
				fmt.Printf("path : %#v, method : %#v, name : %#v\n", route.Path, route.Method, route.Name)
				// routes = append(routes, Route{route.Name, route.Method, route.Path})
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(routesCmd)
}
