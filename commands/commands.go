package commands

import (
	"github.com/urfave/cli"
)

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
