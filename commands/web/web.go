package web

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/meinbaumm/al/commands/config"

	"github.com/urfave/cli/v2"
)

func Web(ctx *cli.Context, urlsToOpen config.CommandConfig) error {
	args := ctx.Args().Slice()
	if len(args) == 0 {
		fmt.Println("Please specify urls to open")
		return nil
	}

	for _, url := range args {
		openUrl(urlsToOpen, url)
	}

	return nil
}

func List(ctx *cli.Context, urlsToShow config.CommandConfig) error {
	verbose := ctx.Bool("v")
	config.ShowCommandConfig(urlsToShow, "urls", verbose)

	return nil
}

func openUrl(urlsToOpen config.CommandConfig, urlName string) {
	if url, ok := urlsToOpen[urlName]; ok {
		err := exec.Command("open", url).Start()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("No such url with name " + urlName)
}
