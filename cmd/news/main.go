package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robtec/newsapi/commands"
	"github.com/urfave/cli"
)

var (
	// Version placeholder for -v option
	Version = "0"
	// CommitId placeholder for -v option
	CommitId = "0"
)

func main() {

	app := cli.NewApp()
	app.Name = "newsapi.org Golang client"
	app.Usage = "a cli for newsapi.org"
	app.UsageText = "news [global options] command [command options] <QUERY>"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key",
			Usage:  "API Key for Authentication",
			EnvVar: "NEWS_API_KEY",
		},
		cli.StringFlag{
			Name:  "url",
			Usage: "Base API url",
			Value: "https://newsapi.org",
		},
	}

	app.Version = fmt.Sprintf("%s - %s", Version, CommitId)

	app.Commands = commands.Commands

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
