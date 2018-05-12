package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

var (
	// ErrMissingAPIKey is returned when an API Key has not been specified
	ErrMissingAPIKey = errors.New("api: API key is missing")
)

// APIKeyHeader is the header key used when sending the API Token
const APIKeyHeader = "X-Api-Key"

// Client used to interact with the NewsAPI
type Client struct {

	// BaseURL the base url of the NewsAPI
	BaseURL *url.URL

	// APIKey is the key used to authenticate against the NewsAPI
	APIKey string

	// httpClient is the client used to communicate with the NewsAPI
	httpClient *http.Client
}

// New returns a configured Client with the supplied http client
// api key and api base url
func New(httpc *http.Client, key string, baseURL string) (*Client, error) {

	if len(key) == 0 {
		return nil, ErrMissingAPIKey
	}

	parsedURL, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	return &Client{parsedURL, key, httpc}, nil
}

func (c *Client) do(req *http.Request) (*Response, error) {

	req.Header.Add(APIKeyHeader, c.APIKey)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	response := new(Response)

	err = json.NewDecoder(resp.Body).Decode(response)

	return response, err
}

// TopHeadlines provides live top and breaking headlines for a country,
// specific category in a country, single source, or multiple sources.
// You can also search with keywords. Articles are sorted by the earliest date published first.
func (c *Client) TopHeadlines(opts Options) (*Response, error) {

	req, err := c.makeRequest("/v2/top-headlines", opts)

	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)

	return resp, err
}

// Everything Search through millions of articles from over 30,000 large and
// small news sources and blogs. This includes breaking news as well as lesser articles.
func (c *Client) Everything(opts Options) (*Response, error) {

	req, err := c.makeRequest("/v2/everything", opts)

	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)

	return resp, err
}

func (c *Client) makeRequest(context string, opts Options) (*http.Request, error) {

	urlValues := optionsToURLValues(opts)

	url := fmt.Sprintf("%s%s?%s", c.BaseURL.String(), context, urlValues.Encode())

	return http.NewRequest("GET", url, nil)
}

func optionsToURLValues(opts Options) url.Values {

	urlValues := url.Values{}

	v := reflect.ValueOf(opts)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i).Name
		value := v.Field(i)

		fieldLower := strings.ToLower(string(field))

		switch value.Kind() {

		case reflect.String:
			if len(value.String()) > 0 {
				urlValues.Add(fieldLower, value.String())
			}

		case reflect.Int64:
			intString := strconv.FormatInt(value.Int(), 10)
			if len(intString) > 0 {
				urlValues.Add(fieldLower, value.String())
			}
		}
	}

	return urlValues
}
