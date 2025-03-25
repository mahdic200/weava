package main

import (
	"fmt"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/cmd"
)

func main() {
	if err := Config.GetEnv(); err != nil {
		fmt.Printf("config, connect : %s\n", err.Error())
	}
	cmd.Execute()
}
