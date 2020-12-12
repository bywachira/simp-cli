package cli

import (
	"github.com/tesh254/simp/helpers"
	"sort"
	"fmt"

	"github.com/urfave/cli/v2"
)

// SetupCLI initialize CLI
func SetupCLI() *cli.App {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "run-commands",
				Aliases: []string{"r"},
				Usage:   "Run your commands on simp.json file",
				Action: func(c *cli.Context) error {
					fmt.Println("We are simping 🤓")
					var cli CLIMethods

					cli.RunAllCommands(helpers.ParseJSONFile().Commands)

					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
