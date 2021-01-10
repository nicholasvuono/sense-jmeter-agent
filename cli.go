package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	flags := []cli.Flag{
		&cli.StringFlag{
			Name: "exec",
		},
		&cli.StringFlag{
			Name: "jmx",
		},
		&cli.StringFlag{
			Name: "jtl",
		},
		&cli.StringFlag{
			Name: "type",
		},
		&cli.StringFlag{
			Name: "url",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "run",
			Usage: "Runs the JMeter test (.jmx) that is specified",
			Flags: flags,
			Action: func(c *cli.Context) error {
				fmt.Println("Running...")
				t := c.String("jmx")
				err := exec.Command(
					c.String("exec"),
					"-n",
					"-t",
					t,
					"-l",
					c.String("jtl"),
				).Run()
				fmt.Println("Complete.")
				if err != nil {
					return err
				}
				res, err := http.Post(
					c.String("url")+"/results/"+c.String("type")+"/add",
					"application/json",
					bytes.NewBuffer(jtltojson.Ptor(c.String("jtl")).JSON()),
				)
				if err != nil {
					return err
				}
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					return (err)
				}
				fmt.Println(string(body))
				return nil
			},
		},
	}

	log.Fatal(app.Run(os.Args))
}
