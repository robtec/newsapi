package commands

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/robtec/newsapi/api"
	"github.com/urfave/cli"
)

func topHeadlines(c *cli.Context) error {

	key := c.GlobalString("key")
	url := c.GlobalString("url")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	client, err := api.New(httpClient, key, url)

	if err != nil {
		return err
	}

	opts := api.Options{Country: "ie"}

	resp, err := client.TopHeadlines(opts)

	fmt.Print(resp)

	return err
}
