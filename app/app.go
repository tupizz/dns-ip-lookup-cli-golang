package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search IPs and name servers on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP's address on the internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:    "servers",
			Aliases: []string{"srvs", "srv"},
			Usage:   "Search name servers on the internet",
			Flags:   flags,
			Action:  searchServers,
		},
	}
	return app
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host) // name server
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
