package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
)

func main() {
	app := fiber.New()

	if err := Config.Connect(); err != nil {
		fmt.Printf("Could not connect to the database\n")
		fmt.Printf("%v\n", err.Error())
		os.Exit(2)
	}

	log.Fatal(app.Listen(":8000"))
}
