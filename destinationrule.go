package main

import (
	"log"
	"github.com/urfave/cli"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DestinationCommand cli.Command
var SpecCommand cli.Command

func init() {
	DestinationCommand = cli.Command{
		Name:    "destination-rule",
		Aliases: []string{"dest"},
		Usage:   "destination create|update|delete|get|list",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list destination rules",
				Action:  destinationList,
			},
			{
				Name:  "get",
				Usage: "get destination rule",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name, n",
						Value: "",
						Usage: "name of destination rule",
					},
				},
				Action: getDestinationRule,
			},
			{
				Name:  "create",
				Usage: "create destination rule ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, n ",
						Value: "",
						Usage: "destination rule",
					},
					cli.StringFlag{
						Name:  "host ",
						Value: "",
						Usage: "destination rule host",
					},
				},
				Subcommands: []cli.Command{
					{
						Name:    "trafficPolicy",
						Aliases: []string{"tp"},
						Usage:   "trafficPolicy ",

						Subcommands: []cli.Command{
							{
								Name:    "portLevelSetting",
								Aliases: []string{"pl"},
								Usage:   "port load strategy ",
								Flags: []cli.Flag{
									cli.StringFlag{
										Name:  "port, p ",
										Value: "",
										Usage: "port number",
									},
								},
								Subcommands: []cli.Command{
									{
										Name:    "loadBalancer",
										Aliases: []string{"lb"},
										Usage:   "load balancer  ",
										Flags: []cli.Flag{
											cli.StringFlag{
												Name:  "type ",
												Value: "simple",
												Usage: "simple ",
											},
											cli.StringFlag{
												Name:  "strategy ",
												Value: "LEAST_CONN",
												Usage: "LEAST_CONN, ROUND_ROBIN etc ",
											},

										},
									},
								},
							},
						},
					},
				},

				Action: createDestinationRule,
			},
			{
				Name:  "delete",
				Usage: "delete destination rule ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name ",
						Value: "",
						Usage: "name of destination rule",
					},
				},
				Action: deleteDestinationRule,
			},
		},
	}

}

func destinationList(c *cli.Context) error {
	ic, namespace, err := setup(c)
	// gateways
	gwList, err := ic.NetworkingV1alpha3().DestinationRules(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get DestinationRule in %s namespace: %s", namespace, err)
	}
	format := c.GlobalString("format")
	if format == "json" {
		fmt.Println(toJsonString(gwList))
	} else if format == "yaml" {
		fmt.Println(toYamlString(gwList))
	} else {

		return cli.NewExitError("Unknown format", 1)

	}
	/*
		for i, rule := range gwList.Items {
			glog.Infof("rule[%d] %s\n\t%s\n", i, rule.Name, rule.APIVersion)
			glog.Infof("  labels:\n")

			for j, label := range rule.Labels {
				glog.Infof("label[%d] %s", j, label)
			}
			glog.Infof("  annotations:\n")
			for j, label := range rule.Annotations {
				glog.Infof("annotation[%d] %s", j, label)
			}
			glog.Infof("  Spec:\n")
			for j, label := range rule.Spec.Host {
				glog.Infof("annotation[%d] %s", j, label)
			}
		}
	*/
	return nil
}
func getDestinationRule(c *cli.Context) error {
	fmt.Printf("get destinationRule %s %+v\n", c.Args().First(), c)
	return nil
}
func createDestinationRule(c *cli.Context) error {
	fmt.Printf("create destinationRule %s %+v\n", c.Args().First(), c)
	return nil
}
func deleteDestinationRule(c *cli.Context) error {
	fmt.Printf("delete destinationRule %s %+v\n", c.Args().First(), c)
	return nil
}
func updateDestinationRule(c *cli.Context) error {
	fmt.Printf("destinationRule %s %+v\n", c.Args().First(), c)
	return nil
}
