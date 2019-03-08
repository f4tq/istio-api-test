package main

import (
	"log"
	"os"
	"github.com/urfave/cli"
	"github.com/golang/glog"
)

func main() {
	glog.V(2)
	app := cli.NewApp()
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "kube-config",
			Value:  "",
			Usage:  "kube config",
			EnvVar: "KUBECONFIG",
		},
		cli.StringFlag{
			Name:  "format, o",
			Value: "yaml",
			Usage: "output format:json or yaml.  default: yaml",
		},
		cli.StringFlag{
			Name:  "namespace, n",
			Value: "default",
			Usage: "name of virtual service",
		},
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "debug toggle",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "verbose toggle",
		},
	}

	app.Commands = []cli.Command{
		DestinationCommand,
		EnvoyFilterCommands,
		GatewayCommands,
		ServiceEntryCommands,
		VirtualServiceCommands,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
