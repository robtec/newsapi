package commands

import (
	"github.com/urfave/cli"
)

// Commands used by the CLI
var Commands = []cli.Command{
	{
		Name:   "th",
		Usage:  "top headlines",
		Action: topHeadlines,
	},
	{
		Name:   "e",
		Usage:  "everything",
		Action: everything,
	},
}
