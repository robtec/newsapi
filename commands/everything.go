package commands

import (
	"crypto/tls"
	"fmt"
	"net/http"

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

	fmt.Print(resp)

	return err
}
