package cli

import (
	cmd "github.schq.secious.com/Logrhythm/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.schq.secious.com/Logrhythm/easy/runner"
)

// Creates the cli structure to be ran from 'main'.
func NewCli() *cmd.App {
	app := cmd.NewApp()
	app.Name = "easy"
	app.Version = "0.0.1"
	app.Usage = "... debug: yeah!!"
	app.Commands = []cmd.Command{
		{
			Name: "plugins",
			Usage: "Finds plugins",
			Action: runner.Run,
			Flags: []cmd.Flag{
				cmd.BoolFlag{
					Name:  "debug",
					Usage: "Outputs the results of finding plugins",
				},
			},
		},
		{
			Name:   "exec",
			Usage:  "Execute * (all) routines",
			Action: runner.Run,
			Flags: []cmd.Flag{
				cmd.StringFlag{
					Name:  "output",
					Value: "./out",
					Usage: "The directory where collected files will be placed (created if it doesn't exist).",
				},
				cmd.StringFlag{
					Name:  "pattern",
					Value: "*",
					Usage: "Run the routines that match the given pattern.",
				},
			},
		},
		{
			Name:  "list",
			Usage: "List all routines in the system.",
			Action: runner.Run,
			Flags: []cmd.Flag{
				cmd.StringFlag{
					Name:  "pattern",
					Value: "*",
					Usage: "Run the routines that match the given pattern.",
				},
			},
		},
	}
	return app
}
