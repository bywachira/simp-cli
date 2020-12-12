package main

import (
	"log"
	"os"
	"github.com/tesh254/simp/cli"
)

func main() {
	app := cli.SetupCLI()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal("Error: ", err)
	}
}
 