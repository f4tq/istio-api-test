package main

import (
	"github.com/urfave/cli"
	"log"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EnvoyFilterCommands cli.Command

func init() {
	EnvoyFilterCommands = cli.Command{
		Name:    "envoy-filter",
		Aliases: []string{"ef"},
		Usage:   "envoy-filter create|update|delete|get|list",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list envoy filters",
				Action:  listEnvoyFilters,
			},
			{
				Name:  "get",
				Usage: "get envoyfilter",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name, n",
						Value: "",
						Usage: "name of envoyfilter",
					},
				},
				Action: getEnvoyFilter,
			},
			{
				Name:   "create",
				Usage:  "create envoyfilter",
				Action: createEnvoyFilter,
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
				Usage:  "delete envoyfilter",
				Action: deleteEnvoyFilter,
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
				Usage:  "update envoyfilter",
				Action: updateEnvoyFilter,
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

func listEnvoyFilters(c *cli.Context) error {
	ic, namespace, err := setup(c)
	// envoy filters
	ruleList, err := ic.NetworkingV1alpha3().EnvoyFilters(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get EnvoyFilters in %s namespace: %s", namespace, err)
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
func createEnvoyFilter(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("create envoy filter %s %+v\n", c.Args().First(), c)
	return nil
}

func getEnvoyFilter(c *cli.Context) error {
	ic, namespace, err := setup(c)
	// envoyfilter
	item, err := ic.NetworkingV1alpha3().EnvoyFilters(namespace).Get(c.String("name"), metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Failed to get EnvoyFilter in %s namespace: %s", namespace, err)
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
func deleteEnvoyFilter(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("delete envoy filter %s %+v\n", c.Args().First(), c)
	return nil
}
func updateEnvoyFilter(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("update envoy filter %s %+v\n", c.Args().First(), c)
	return nil
}
