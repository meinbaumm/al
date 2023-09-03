package open

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/meinbaumm/al/commands/config"

	"github.com/urfave/cli/v2"
)

func Open(ctx *cli.Context, appsToOpen config.CommandConfig) error {
	args := ctx.Args().Slice()
	if len(args) == 0 {
		fmt.Println("Please specify apps to open")
		return nil
	}

	for _, app := range args {
		openApp(appsToOpen, app)
	}

	return nil
}

func openApp(appsToOpen config.CommandConfig, appName string) {
	if app, ok := appsToOpen[appName]; ok {
		err := exec.Command("open", app).Start()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("No such application with name " + appName)
}

func List(ctx *cli.Context, appsToOpen config.CommandConfig) error {
	verbose := ctx.Bool("v")
	config.ShowCommandConfig(appsToOpen, "apps", verbose)

	return nil
}
