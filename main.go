package kubec

import (
	"fmt"
	"os"

	"github.com/jedipunkz/kubec/commands"
	"github.com/mitchellh/cli"
)

const (
	app     = "kubec"
	version = "0.0.1"
)

func main() {
	c := cli.NewCLI(app, version)
	c.Args = os.Args[1:]
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c.Commands = map[string]cli.CommandFactory{
		"list deployments": func() (cli.Command, error) {
			return &commands.ListDeployments{UI: &cli.ColoredUi{Ui: ui, WarnColor: cli.UiColorYellow, ErrorColor: cli.UiColorRed}}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		ui.Error(fmt.Sprintf("Error: %s", err))
	}

	os.Exit(exitStatus)
}
