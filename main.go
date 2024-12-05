package main

import (
	"log"
	"os"

	"github.com/k-nox/advent-of-code-2024/cli"
)

func main() {
	app := cli.App()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
