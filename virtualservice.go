package main

import (
	"github.com/urfave/cli"
	"github.com/golang/glog"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VirtualServiceCommands cli.Command

func init() {
	VirtualServiceCommands = cli.Command{
		Name:
		"virtual-service",
		Aliases: []string{"vs"},
		Usage:   "virtual-service create|update|delete|get|list",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list virtual services",
				Action:  listVirtualService,
			},
			{
				Name:  "get",
				Usage: "get virtual service",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name, n",
						Value: "",
						Usage: "name of virtual service",
					},
				},
				Action: getVirtualService,
			},

			{
				Name:  "create",
				Usage: "create a virtual service",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
						Usage: "file to load description from",
					},
				},
				Action: createVirtualService,
			},
			{
				Name:  "delete",
				Usage: "delete a virtual service",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
						Usage: "file to load description from",
					},
				},
				Action: deleteVirtualService,
			},
			{
				Name:  "update",
				Usage: "update a virtual service",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: "",
						Usage: "file to load description from",
					},
				},
				Action: updateVirtualService,
			},
		},
	}
}

func listVirtualService(c *cli.Context) error {
	ic, namespace, err := setup(c)
	if err != nil {
		return err
	}
	vsList, err := ic.NetworkingV1alpha3().VirtualServices(namespace).List(metav1.ListOptions{})
	if err != nil {
		glog.Fatalf("Failed to get VirtualService in %s namespace: %s", namespace, err)
	}
	format := c.GlobalString("format")

	if format == "json" {
		fmt.Println(toJsonString(vsList))
	} else if format == "yaml" {
		fmt.Println(toYamlString(vsList))
	} else {

		return cli.NewExitError("Unknown format", 1)

	}
	return nil
	/* for i := range vsList.Items {
			vs := vsList.Items[i]
			glog.Infof("Index: %d VirtualService Hosts: %+v\n", i, vs.Spec.GetHosts())
		}
	*/
	return nil
}
func getVirtualService(c *cli.Context) error {
	ic, namespace, err := setup(c)
	if err != nil {
		return err
	}
	vsList, err := ic.NetworkingV1alpha3().VirtualServices(namespace).List(metav1.ListOptions{LabelSelector: namespace, FieldSelector: c.String("name")})
	for i := range vsList.Items {
		vs := vsList.Items[i]
		glog.Infof("Index: %d VirtualService Hosts: %+v\n", i, vs.Spec.Hosts)
	}
	return nil
}
func createVirtualService(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("create virtual service %s %+v\n", c.Args().First(), c)
	return nil
}
func deleteVirtualService(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("delete virtual service %s %+v\n", c.Args().First(), c)
	return nil
}
func updateVirtualService(c *cli.Context) error {
	if c.GlobalIsSet("debug") {

	}
	fmt.Printf("update virtual service %s %+v\n", c.Args().First(), c)
	return nil
}
