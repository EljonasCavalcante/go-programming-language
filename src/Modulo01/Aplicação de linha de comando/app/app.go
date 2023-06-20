package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronto para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = " Aplicação de linha de Comando"
	app.Usage = "Busca IPs e Nomes de servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "teste.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereço de internet",
			Flags:  flags,
			Action: BuscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na Internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func BuscarIps(c *cli.Context) {
	host := c.String("host")

	//usar pacote net

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	// observe que usou o _ , para recebe apenas uma resposta
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) // nome server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
