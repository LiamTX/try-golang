package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Generate returns the command line app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "command line app"
	app.Usage = "search ips and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "search ips",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "search server names",
			Flags:  flags,
			Action: searchServerNames,
		},
	}

	return app
}

func searchServerNames(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
