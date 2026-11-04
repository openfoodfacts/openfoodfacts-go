// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

// This is a go library used to access the OpenFoodFacts.org database for food product, ingredients and
// nutritional data from within your go application.
//
// The main method of using this library is to create a DataOperator and call methods on it.
package openfoodfacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const defaultUserAgent = "OpenFoodFacts - Go - v0.0 - https://github.com/openfoodfacts/openfoodfacts-go"

var (
	// ErrNoProduct is an error returned by Client.Product when the product could not be
	// retrieved successfully.
	// It is not a transient error, the product does not exist.
	ErrNoProduct = errors.New("Product retrieval failure")

	// ErrUnauthorized is an error returned by Client methods that require a valid user account, but none
	// was provided when the Client was instantiated.
	ErrUnauthorized = errors.New("Action requires user account")
)

// Option is a functional option type that allows us to configure the Client.
type Option func(*Client)

// Client is an OpenFoodFacts client.
// It uses the official API as data source.
type Client struct {
	locale    string
	username  string
	password  string
	live      bool
	url       url.URL
	client    *http.Client
	userAgent string
}

// NewClient returns a Client that is capable of talking to the official OpenFoodFacts database via
// the HTTP API, or the dev server if live is false.
//
// The locale should be one of "world" or the country level code for the locale you wish to use.
//
// The username and password should be set to your OpenFoodFacts credentials if you need WRITE access, else provide
// them both as empty strings for anonymous access.
//
// # Sandbox mode
//
// If you are testing your application, you should use the test server in order to use the sandbox environment instead
// of the live servers. See the Sandbox() method for more detail and an example.
//
// # Timeout
//
// By default the HTTP client doesn't sets a timeout. See the Timeout(time.Duration) method for more detail.
//
// # UserAgent
//
// Please set a UserAgent HTTP Header with the name of the app/service querying, the version, system and a URL if
// you have one, so that you are not blocked by mistake
// (e.g. CoolFoodApp - Go - Version 1.0 - https://coolfoodapp.com)
func NewClient(opts ...Option) Client {
	client := Client{
		locale:    "world",
		client:    &http.Client{},
		live:      true,
		username:  "",
		password:  "",
		userAgent: defaultUserAgent,
	}

	client.Apply(opts...)

	return client
}

func (c *Client) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	c.url = buildURL(c)
}

func WithLocale(locale string) Option {
	return func(c *Client) {
		c.locale = locale
	}
}

func WithCredentials(username, password string) Option {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

func WithSandbox() Option {
	return func(c *Client) {
		c.live = false
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.client = &http.Client{Timeout: timeout}
	}
}

func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// Product returns a new Product for the given code, retrieved from the server.
//
// It will return an error on a failed retrieval, if the retrieval is successful but the API result status is not 1,
// then will return a "ProductRetrievalError" error. This indicates the product is not available.
func (h *Client) Product(code string) (*Product, error) {
	request := h.newRequest("GET", "/api/v0/product/%s.json", code)

	resp, err := h.client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	productResult := &ProductResult{}
	err = json.Unmarshal(body, productResult)

	if err != nil {
		offs := int64(-1)
		if e, ok := err.(*json.UnmarshalTypeError); ok {
			offs = e.Offset
		} else if e, ok := err.(*json.SyntaxError); ok {
			offs = e.Offset
		}

		if offs > -1 {
			a := offs - 50
			b := offs + 20
			n := int64(len(body))
			if a < 0 {
				a = 0
			}
			if b > n {
				b = n
			}
			err = fmt.Errorf("%s at:\n  %s⚠️ %s",
				string(err.Error()),
				string(body[a:offs]), string(body[offs:b]),
			)
		}

		return nil, err
	}

	if productResult.Status != 1 {
		return nil, ErrNoProduct
	}

	return productResult.Product, nil
}

func buildURL(c *Client) url.URL {
	sub, tld := "ssl-api", "org"
	if !c.live {
		sub, tld = "world", "net"
	}
	if c.locale != "world" {
		sub = c.locale
	}
	return url.URL{Scheme: "https", Host: fmt.Sprintf("%s.openfoodfacts.%s", sub, tld)}
}

// newRequest is an internal function to setup the request based on the given
// locale/liveness of the given Client.
func (h *Client) newRequest(method, format string, args ...interface{}) *http.Request {
	u := h.url
	u.Path = fmt.Sprintf(format, args...)
	request, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil
	}

	if !h.live {
		request.SetBasicAuth("off", "off")
	}

	if h.userAgent != "" {
		request.Header.Set("User-Agent", h.userAgent)
	}

	return request
}
