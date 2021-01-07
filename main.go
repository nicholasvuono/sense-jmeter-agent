package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	log.Fatal(cli.NewApp().Run(os.Args))
}
