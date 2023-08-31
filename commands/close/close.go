package close

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/meinbaumm/al/commands"

	"github.com/urfave/cli/v2"
)

type appsMap map[string]string

func Close(ctx *cli.Context, appsToClose appsMap) error {
	args := ctx.Args().Slice()
	if len(args) == 0 {
		fmt.Println("Please specify apps to close")
		return nil
	}

	for _, app := range args {
		if app == "showall" {
			showAllApps(appsToClose)
			continue
		}

		closeApp(appsToClose, app)
	}

	return nil
}

func closeApp(appsToClose appsMap, appName string) {
	if app, ok := appsToClose[appName]; ok {
		err := exec.Command("osascript", "-e", app).Run()
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
