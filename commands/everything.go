package commands

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/robtec/newsapi/api"
	"github.com/urfave/cli"
)

func everything(c *cli.Context) error {

	key := c.GlobalString("key")
	url := c.GlobalString("url")

	query := c.Args().Get(0)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	client, err := api.New(httpClient, key, url)

	if err != nil {
		return err
	}

	opts := api.Options{Q: query}

	resp, err := client.Everything(opts)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "Source"})
	table.SetRowLine(true)

	for _, v := range resp.Articles {
		table.Append([]string{v.Title, v.Source.Name})
	}

	table.Render()

	return err
}
