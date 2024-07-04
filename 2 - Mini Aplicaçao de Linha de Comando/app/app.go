package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicaçao
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicaçao de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPS de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.b",
				},
			},
			Action: buscarIpts,
		},
	}

	return app
}

func buscarIpts(c *cli.Context) {
	host := c.String("host")

	// net
	ips, errro := net.LookupIP(host)
	if errro != nil {
		log.Fatal(errro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
