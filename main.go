package al

import (
	"fmt"
	"log"
	"os"

	"github.com/meinbaumm/al/commands/close"
	"github.com/meinbaumm/al/commands/open"
	"github.com/meinbaumm/al/commands/web"

	"github.com/urfave/cli/v2"
)

func main() {
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading config file: %w", err))
	}

	app := cli.NewApp()
	app.Name = "al"
	app.Usage = "Open/close web urls/apps"
	app.Version = "0.1.1"
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
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
