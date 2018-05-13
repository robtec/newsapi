package commands

import (
	"time"

	humanize "github.com/dustin/go-humanize"
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
			},
		},
		Action: topHeadlines,
	},
	{
		Name:  "everything",
		Usage: "Everything",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "l",
				Usage: "Language",
				Value: "en",
			},
			cli.StringFlag{
				Name:  "s",
				Usage: "SortBy",
				Value: "popularity",
			},
		},
		Action: everything,
	},
}

func prettyTime(ugly string) string {
	parsed, _ := time.Parse(time.RFC3339, ugly)
	return humanize.Time(parsed)
}
