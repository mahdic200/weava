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
		type SingleRoute struct {
			Name   string
			Path   string
			Method string
		}
		routes := make([]SingleRoute, 0)
	biggerLoop:
		for _, route := range Config.App.GetRoutes() {
			if route.Path != "/*" && route.Name != "" {
				for _, i_route := range routes {
					if i_route.Name == route.Name || (strings.Contains(route.Name, "index") && route.Method == "HEAD") {
						continue biggerLoop
					}
				}
				routes = append(routes, SingleRoute{Name: route.Name, Path: route.Path, Method: route.Method})
			}
		}
		for _, route := range routes {
			fmt.Printf("%v\n", route)
		}
	},
}

func init() {
	rootCmd.AddCommand(routesCmd)
}
