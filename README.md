# NewsAPI Go

[![Go Report Card](https://goreportcard.com/badge/github.com/robtec/newsapi-go)](https://goreportcard.com/report/github.com/robtec/newsapi-go) [![GoDoc](https://godoc.org/github.com/robtec/newsapi/api?status.svg)](https://godoc.org/github.com/robtec/newsapi/api)

Golang client and library for https://newsapi.org

## Install

```bash
$ go get github.com/robtec/newsapi

$ go install ./cmd/news
```

## CLI Usage

You must register with News API to get an API Key

https://newsapi.org/register


```bash
$ export NEWS_API_KEY=<YOUR-KEY>

# Get the Top Headlines from Ireland, with 'Eurovision' as the Query
$ news top -c ie "Eurovision"

+--------------------------------+--------------------------------+----------------+-------------+
|             TITLE              |          DESCRIPTION           |     SOURCE     |   POSTED    |
+--------------------------------+--------------------------------+----------------+-------------+
| How Ryan's Eurovision chutzpah | Didn't the heart nearly burst  | Independent.ie | 2 hours ago |
| outfoxed us all                | out of my chest when the       |                |             |
|                                | Portuguese gals chimed in      |                |             |
|                                | unison                         |                |             |
+--------------------------------+--------------------------------+----------------+-------------+
...

```

See `news help` for more options


## Library Usage

```golang
package main

import (

	"fmt"
	"net/http"

	"github.com/robtec/newsapi/api"
)

func main() {

    httpClient := http.Client{}
    key := "my-api-key"
    url := "https://newsapi.org"

    query := "Elon Musk"

    // Create a client, passing in the above
    client, err := api.New(httpClient, key, url)

    // Create options for Ireland and Business
    opts := api.Options{Country: "ie", Category: "business"}

    // Get Top Headlines with options from above
    topHeadlines, err := client.TopHeadlines(opts)

    // Different options
    moreOpts := api.Options{Language: "en", Q: query, SortBy: "popularity"}

     // Get Everything with options from above
    everything, err := client.Everything(moreOpts)

}
```

## Contribution

Checkout source

```bash

$ mkdir -p $GOPATH/src/github.com/<your-username> && cd $_
$ git clone https://github.com/robtec/newsapi-go.git

$ cd newsapi

$ make deps // install dependencies
```

Checkout the `Makefile` for extra development options