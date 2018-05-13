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

	language := c.String("l")
	sortBy := c.String("s")

	query := c.Args().Get(0)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	client, err := api.New(httpClient, key, url)

	if err != nil {
		return err
	}

	opts := api.Options{Q: query, Language: language, SortBy: sortBy}

	resp, err := client.Everything(opts)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "Description", "Source", "Posted"})
	table.SetRowLine(true)

	for _, v := range resp.Articles {
		prettyTime := prettyTime(v.PublishedAt)
		table.Append([]string{v.Title, v.Description, v.Source.Name, prettyTime})
	}

	table.Render()

	return err
}
