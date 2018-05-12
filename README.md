# NewsAPI Go

Golang client and library for https://newsapi.org

## Install

```bash
go get github.com/robtec/newsapi

go install ./cmd/news
```

## CLI Usage

```bash
$ news -h

NAME:
   newsapi.org Golang client - a cli for newsapi.org

USAGE:
   news [global options] command [command options] [arguments...]

VERSION:
   0.1 - 

COMMANDS:
     th       top headlines
     e        everything
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --key value    API Key for Authentication [$NEWS_API_KEY]
   --url value    Base API url (default: "https://newsapi.org")
   --help, -h     show help
   --version, -v  print the version
```


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

    // Create a client, passing in the above
    client, err := api.New(httpClient, key, url)

    // Create options for Ireland and Business
    opts := api.Options{Country: "ie", Category: "business"}

    resp, err := client.TopHeadlines(opts)

    fmt.Print(resp)
}
```

## Contribution

Checkout source

```bash

$ mkdir -p $GOPATH/src/github.com/<your-username> && cd $_
$ git clone https://github.com/robtec/newsapi-go.git
$ cd newsapi-go

$ make deps // install dependencies
```

Checkout the `Makefile` for extra development options