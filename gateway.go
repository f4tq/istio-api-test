package main

import (
	"github.com/urfave/cli"
	"log"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

var GatewayCommands cli.Command

func init(){
	GatewayCommands = cli.Command     {
		Name:    "gateway",
		Aliases: []string{"gw"},
		Usage:   "gateway create|update|delete|get|update",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Aliases: []string{"ls"},
				Usage: "list gateways",
				Action: listGateways,
			},
			{
				Name:  "get",
				Usage: "get gateway",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name, n",
						Value: "",
						Usage: "name of gateway",
					},
				},
				Action: getGateway,
			},
			{
				Name:  "create",
				Usage: "create gateway",
				Action: createGateway,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
						Usage: "file to load description from",
					},
				},
			},
			{
				Name:  "delete",
				Usage: "delete gateway",
				Action: deleteGateway,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
						Usage: "file to load description from",
					},
				},
			},
			{
				Name:  "update",
				Usage: "update gateway",
				Action: updateGateway,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
						Usage: "file to load description from",
					},
				},
			},
		},
	}
}
func listGateways (c *cli.Context) error {
	ic,namespace,err := setup(c)
	// gateways
	ruleList, err := ic.NetworkingV1alpha3().Gateways(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get Gateways in %s namespace: %s", namespace, err)
	}
	format := c.GlobalString("format")
	if format == "json" {
		fmt.Println(toJsonString(ruleList))
	} else if format == "yaml" {
			fmt.Println(toYamlString(ruleList))
	} else {

		return cli.NewExitError("Unknown format",1)

	}
	return nil
	/*
	for i,rule := range ruleList.Items {
		//rule := ruleList.Items[i]
		log.Printf("Index: %d Rule name %s / ns %s : %+v\n", i, rule.Name, rule.Spec.Servers, rule.Spec)
	}
	*/
}
func getGateway(c *cli.Context) error {
	//debug:= c.GlobalIsSet("debug")
	name:= c.GlobalString("name")
	ic,namespace,err := setup(c)
	// gateways
	ruleList, err := ic.NetworkingV1alpha3().Gateways(namespace).Get(name,metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Failed to get Gateways in %s namespace: %s", namespace, err)
	}
	format := c.GlobalString("format")
	if format == "json" {
		fmt.Println(toJsonString(ruleList))
	} else if format == "yaml" {
		fmt.Println(toYamlString(ruleList))
	} else {
		return cli.NewExitError("Unknown format",1)
	}
	return nil
}

func createGateway(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("create gateway %s %+v\n", c.Args().First(), c)
	return nil
}
func deleteGateway(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("delete gateway %s %+v\n", c.Args().First(), c)
	return nil
}
func updateGateway(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("update gateway %s %+v\n", c.Args().First(), c)
	return nil
}