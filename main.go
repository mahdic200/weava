package main

import (
	"fmt"
	"os"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Routes"
	"github.com/mahdic200/weava/cmd"
)

func main() {
	/* These two if sections are so important, as much as if they fail ,
	application must exit */
	if err := Config.GetEnv(); err != nil {
		fmt.Printf("Config, connect : %s\n", err.Error())
		os.Exit(2)
	}
	if err := Config.Connect(); err != nil {
		fmt.Printf("Could not connect to the database\n")
		fmt.Printf("%v\n", err.Error())
		os.Exit(2)
	}
	Routes.SetupRoutes(Config.App)
	cmd.Execute()
}
