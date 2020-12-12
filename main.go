package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ssler",
		Usage: "create self sign SSL with ease.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "country",
				Aliases: []string{"c"},
				Value:   "NL",
				Usage:   "Organization Country",
			},
			&cli.StringFlag{
				Name:    "state",
				Aliases: []string{"st"},
				Value:   "North Holland",
				Usage:   "Organization State",
			},
			&cli.StringFlag{
				Name:    "city",
				Aliases: []string{"l"},
				Value:   "Amsterdam",
				Usage:   "Organization City",
			},
			&cli.StringFlag{
				Name:    "common-name",
				Aliases: []string{"cn"},
				Value:   "Liandm CA",
				Usage:   "Common Name",
			},
			&cli.StringFlag{
				Name:    "organization",
				Aliases: []string{"o"},
				Value:   "Liandm, Ltd.",
				Usage:   "Organization Name",
			},
			&cli.StringFlag{
				Name:    "organization-unit",
				Aliases: []string{"ou"},
				Value:   "IT Department",
				Usage:   "Organization Unit",
			},
			&cli.StringFlag{
				Name:    "ca-path",
				Aliases: []string{"cap"},
				Value:   "",
				Usage:   "CA file path",
			},
			&cli.StringFlag{
				Name:    "certification-path",
				Aliases: []string{"cp"},
				Value:   "",
				Usage:   "Certification file path",
			},
			&cli.StringFlag{
				Name:    "domains",
				Aliases: []string{"d"},
				Value:   "localhost,127.0.0.1,::1,example.test,api.example.dev",
				Usage:   "Comma separate domain names",
			},
		},
		Action: func(c *cli.Context) error {
			name := "create"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			country := c.String("country")
			state := c.String("state")
			city := c.String("city")
			commonName := c.String("common-name")
			organization := c.String("organization")
			organizationUnit := c.String("organization-unit")
			caPath := c.String("ca-path")
			certificationPath := c.String("certification-path")

			cmdString := "openssl req -x509 -nodes -new -sha512 -days 365 -newkey rsa:4096 -keyout ca.key -out ca.pem -subj \"/C=" + country + "/CN=" + commonName + "/O=" + organization + "\""
			cmdString += " && openssl x509 -in " + caPath + "ca.pem -text -noout"
			cmdString += " && openssl x509 -outform pem -in " + caPath + "ca.pem -out " + caPath + "ca.crt"
			cmdString += " && openssl req -new -nodes -newkey rsa:4096 -keyout " + certificationPath + "localhost.key -out " + certificationPath + "localhost.csr -subj \"/C=" + country + "/ST=" + state + "/L=" + city + "/O=" + organization + "/OU=" + organizationUnit + "/CN=localhost\""
			cmdString += " && openssl x509 -req -sha512 -days 365 -extfile v3.ext -CA " + caPath + "ca.crt -CAkey " + caPath + "ca.key -CAcreateserial -in " + certificationPath + "localhost.csr -out " + certificationPath + "localhost.crt"

			cmd := exec.Command(cmdString)
			err := cmd.Run()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(name)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
