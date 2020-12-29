package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "testapp",
		Usage: "urfave/cli test commandline app",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello, CLI!")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "load configuration from `FILE`",
				EnvVars: []string{"CONFIG_FILE"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "subcmd",
				Usage: "urfave/cli test subcommandline app",
				Action: func(c *cli.Context) error {
					fmt.Println("Hello, CLI subcommand")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
