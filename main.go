package main

import (
	"fmt"
	"log"
	"os"

	"github.com/meinbaumm/al/commands/close"
	"github.com/meinbaumm/al/commands/open"
	"github.com/meinbaumm/al/commands/web"
	"github.com/meinbaumm/al/config"

	"github.com/urfave/cli/v2"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading config file: %w", err))
	}

	app := cli.NewApp()
	app.Name = "al"
	app.Usage = "Open/close web urls/apps"
	app.Version = "0.1.3"
	app.Authors = []*cli.Author{
		{
			Name:  "Maxim Petrenko",
			Email: "meinbaumm@gmail.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "web",
			Usage: "Open web urls",
			Action: func(ctx *cli.Context) error {
				err := web.Web(ctx, cfg.Urls)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:  "list",
					Usage: "List all web urls in config file",
					Action: func(ctx *cli.Context) error {
						err := web.List(ctx, cfg.Urls)
						if err != nil {
							log.Fatal(err)
						}

						return nil
					},
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "verbose",
							Aliases: []string{"v"},
							Usage:   "Show keys and values",
						},
					},
				},
			},
		},
		{
			Name:  "open",
			Usage: "Open apps",
			Action: func(ctx *cli.Context) error {
				err := open.Open(ctx, cfg.Open)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:  "list",
					Usage: "List all apps-to-open in config file",
					Action: func(ctx *cli.Context) error {
						err := open.List(ctx, cfg.Urls)
						if err != nil {
							log.Fatal(err)
						}

						return nil
					},
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "verbose",
							Aliases: []string{"v"},
							Usage:   "Show keys and values",
						},
					},
				},
			},
		},
		{
			Name:  "close",
			Usage: "Close apps",
			Action: func(ctx *cli.Context) error {
				err := close.Close(ctx, cfg.Close)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:  "list",
					Usage: "List all apps-to-close in config file",
					Action: func(ctx *cli.Context) error {
						err := open.List(ctx, cfg.Urls)
						if err != nil {
							log.Fatal(err)
						}

						return nil
					},
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "verbose",
							Aliases: []string{"v"},
							Usage:   "Show keys and values",
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
