package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"

	jtltojson "github.com/nicholasvuono/jtl-to-json"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sense JMeter Agent"
	app.Usage = "This is a CLI built using Go to allow users to run a JMeter test and send the results to an implemetation of the sense-microservice to be formatted and stored!"
	app.Version = "1.0.0"

	app.Commands = []*cli.Command{
		{
			Name:  "run",
			Usage: "Runs the JMeter test (.jmx) that is specified",
			Action: func(c *cli.Context) error {
				err := exec.Command(
					c.String("exec"),
					"-n",
					"-t "+c.String("jmx"),
					"-l "+c.String("jtl"),
				).Run()
				if err != nil {
					return err
				}
				_, err = http.Post(
					c.String("url")+"/results/"+c.String("type")+"/add",
					"application/json",
					bytes.NewBuffer(jtltojson.Ptor(c.String("jtl")).JSON()),
				)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	log.Fatal(app.Run(os.Args))
}
