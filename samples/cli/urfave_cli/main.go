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
		Before: func(c *cli.Context) error {
			// 共通の事前処理（例: 設定読み込み）
			fmt.Println("before: setup")
			return nil
		},
		After: func(c *cli.Context) error {
			// 共通の後処理（例: ログ出力/終了処理）
			fmt.Println("after: cleanup")
			return nil
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Hello, CLI!")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "load configuration from `FILE`",
				// 未指定時のデフォルトを用意しておく
				Value:   "./config.yaml",
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

// 実行例:
//   go run main.go --help
//   go run main.go subcmd --help
