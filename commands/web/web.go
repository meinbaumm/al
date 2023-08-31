package web

import (
	"fmt"
	"log"
	"os/exec"

	"al/commands"

	"github.com/urfave/cli/v2"
)

type urlsMap map[string]string

func Web(ctx *cli.Context, urlsToOpen urlsMap) error {
	args := ctx.Args().Slice()
	if len(args) == 0 {
		fmt.Println("Please specify urls to open")
		return nil
	}

	for _, url := range args {
		if url == "showall" {
			showAllURLs(urlsToOpen)
			continue
		}

		openUrl(urlsToOpen, url)
	}

	return nil
}

func openUrl(urlsToOpen urlsMap, urlName string) {
	if url, ok := urlsToOpen[urlName]; ok {
		err := exec.Command("open", url).Start()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("No such url with name " + urlName)
}

func showAllURLs(urlsToShow urlsMap) {
	commands.Show(commands.GetSortedMapKeys(urlsToShow), "urls")
}
