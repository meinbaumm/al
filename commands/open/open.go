package open

import (
	"fmt"
	"log"
	"os/exec"

	"al/commands"

	"github.com/urfave/cli/v2"
)

type appsMap map[string]string

func Open(ctx *cli.Context, appsToOpen appsMap) error {
	args := ctx.Args().Slice()
	if len(args) == 0 {
		fmt.Println("Please specify apps to open")
		return nil
	}

	for _, app := range args {
		if app == "showall" {
			showAllApps(appsToOpen)
			continue
		}

		openApp(appsToOpen, app)
	}

	return nil
}

func openApp(appsToOpen appsMap, appName string) {
	if app, ok := appsToOpen[appName]; ok {
		err := exec.Command("open", app).Start()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("No such application with name " + appName)
}

func showAllApps(appsToShow appsMap) {
	commands.Show(commands.GetSortedMapKeys(appsToShow), "apps")
}
