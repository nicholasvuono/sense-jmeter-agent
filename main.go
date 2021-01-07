package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sense JMeter Agent"
	app.Usage = "This is a CLI built using Go to allow users to run a JMeter test and send the results to an implemetation of the sense-microservice to be formatted and stored!"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Runs the JMeter test (.jmx) that is specified",
			Action: func(c *cli.Context) error {
				_, err := exec.Command("")
			},
		},
	},
		log.Fatal(app.Run(os.Args))
}
