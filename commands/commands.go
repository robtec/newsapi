package commands

import (
	"github.com/urfave/cli"
)

// Commands used by the CLI
var Commands = []cli.Command{
	{
		Name:  "top",
		Usage: "top headlines",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "c",
				Usage: "Country",
			}},
		Action: topHeadlines,
	},
	{
		Name:   "everything",
		Usage:  "everything",
		Action: everything,
	},
}
