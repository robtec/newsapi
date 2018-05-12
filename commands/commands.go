package commands

import (
	"github.com/urfave/cli"
)

// Commands used by the CLI
var Commands = []cli.Command{
	{
		Name:   "top",
		Usage:  "top headlines",
		Action: topHeadlines,
	},
	{
		Name:   "everything",
		Usage:  "everything",
		Action: everything,
	},
}
