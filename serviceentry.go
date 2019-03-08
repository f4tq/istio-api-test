package main

import (
	"github.com/urfave/cli"
	"log"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ServiceEntryCommands cli.Command

func init() {
	ServiceEntryCommands = cli.Command{
		Name:    "service-entry",
		Aliases: []string{"se"},
		Usage:   "service-entry create|update|delete|get|list",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list serviceentries",
				Action:  listServiceEntries,
			},
			{
				Name:  "get",
				Usage: "get serviceentry",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name, n",
						Value: "",
						Usage: "name of serviceentry",
					},
				},
				Action: getServiceEntry,
			},
			{
				Name:   "create",
				Usage:  "create serviceentry",
				Action: createServiceEntry,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
						Usage: "file to load description from",
					},
				},
			},
			{
				Name:   "delete",
				Usage:  "delete serviceentry",
				Action: deleteServiceEntry,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
						Usage: "file to load description from",
					},
				},
			},
			{
				Name:   "update",
				Usage:  "update serviceentry",
				Action: updateServiceEntry,
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
func listServiceEntries(c *cli.Context) error {
	ic, namespace, err := setup(c)
	// serviceentries
	ruleList, err := ic.NetworkingV1alpha3().ServiceEntries(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get ServiceEntries in %s namespace: %s", namespace, err)
	}
	format := c.GlobalString("format")
	if format == "json" {
		fmt.Println(toJsonString(ruleList))
	} else if format == "yaml" {
		fmt.Println(toYamlString(ruleList))
	} else {

		return cli.NewExitError("Unknown format", 1)

	}
	return nil
	/*
	for i,rule := range ruleList.Items {
		//rule := ruleList.Items[i]
		log.Printf("Index: %d Rule name %s / ns %s : %+v\n", i, rule.Name, rule.Spec.Servers, rule.Spec)
	}
	*/
}
func getServiceEntry(c *cli.Context) error {
	ic, namespace, err := setup(c)
	// serviceentries
	item, err := ic.NetworkingV1alpha3().ServiceEntries(namespace).Get(c.String("name"), metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Failed to get ServiceEntries in %s namespace: %s", namespace, err)
	}
	format := c.GlobalString("format")
	if format == "json" {
		fmt.Println(toJsonString(item))
	} else if format == "yaml" {
		fmt.Println(toYamlString(item))
	} else {

		return cli.NewExitError("Unknown format", 1)

	}
	return nil
}

func createServiceEntry(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("create service entry %s %+v\n", c.Args().First(), c)
	return nil
}
func deleteServiceEntry(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("delete service entry %s %+v\n", c.Args().First(), c)
	return nil
}
func updateServiceEntry(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("update service entry %s %+v\n", c.Args().First(), c)
	return nil
}
