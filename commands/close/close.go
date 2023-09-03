package close

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/meinbaumm/al/config"

	"github.com/urfave/cli/v2"
)

func Close(ctx *cli.Context, appsToClose config.CommandConfig) error {
	args := ctx.Args().Slice()
	if len(args) == 0 {
		fmt.Println("Please specify apps to close")
		return nil
	}

	for _, app := range args {
		closeApp(appsToClose, app)
	}

	return nil
}

func closeApp(appsToClose config.CommandConfig, appName string) {
	if app, ok := appsToClose[appName]; ok {
		err := exec.Command("osascript", "-e", app).Run()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("No such application with name " + appName)
}

func List(ctx *cli.Context, appsToClose config.CommandConfig) error {
	verbose := ctx.Bool("v")
	config.ShowCommandConfig(appsToClose, "apps", verbose)

	return nil
}
