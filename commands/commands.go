package commands

import (
	"github.com/urfave/cli"
)

// Commands used by the CLI
var Commands = []cli.Command{
	{
		Name:  "top",
		Usage: "Top Headlines",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "c",
				Usage: "Country",
				Value: "ie",
			}},
		Action: topHeadlines,
	},
	{
		Name:   "everything",
		Usage:  "Everything",
		Action: everything,
	},
}
