// Copyright © 2016 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HttpApi is a DataOperator that uses the official API for it's data source.
type HttpApi struct {
	locale   string
	username string
	password string
	live     bool
}

// NewHttpApiOperator returns a DataOperator that is capable of talking to the official OpenFoodFacts database via
// the HTTP API, or the dev server if live is false.
//
// The locale should be one of "world" or the country level code for the locale you wish to use.
//
// The username and password should be set to your OpenFoodFacts credentials if you need WRITE access, else provide
// them both as empty strings for anonymous access.
//
// Sandbox mode
//
// If you are testing your application, you should use the test server in order to use the sandbox environment instead
// of the live servers. See the Sandbox() method for more detail and an example.
func NewHttpApiOperator(locale, username, password string) DataOperator {
	return &HttpApi{
		locale:   locale,
		username: username,
		password: password,
		live:     true,
	}
}

// GetProduct returns a new Product for the given code, retrieved from the server.
//
// It will return an error on a failed retrieval, if the retrieval is successful but the API result status is not 1,
// then will return a "ProductRetrievalError" error. This indicates the product is not available.
func (h *HttpApi) GetProduct(code string) (*Product, error) {
	request := h.newRequest("GET", "/api/v0/product/%s.json", code)

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

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
			err = errors.New(
				fmt.Sprintf("%s at:\n  %s⚠️ %s",
					string(err.Error()),
					string(body[a:offs]), string(body[offs:b]),
				),
			)
		}

		return nil, err
	}

	if productResult.Status != 1 {
		return nil, ProductRetrievalError
	}

	return productResult.Product, nil
}

// Sandbox configures this operator to use the sandbox server at http://world.openfoodfacts.net instead of the live
// server. This is used for testing purposes instead of operating on the live server.
func (h *HttpApi) Sandbox() {
	h.live = false
}

// newRequest is an internal function to setup the request based on the given locale/liveness of the given HttpApi.
func (h *HttpApi) newRequest(method, format string, args ...interface{}) *http.Request {
	path := fmt.Sprintf(format, args...)
	scheme := "https"
	sub := "ssl-api"
	tld := "org"

	if !h.live {
		scheme = "http"
		sub = "world"
		tld = "net"
	}

	if h.locale != "world" {
		// Currently there is no way to set the locale but by hitting the non-https locale specific sub-domain.
		// See: https://github.com/openfoodfacts/openfoodfacts-server/issues/573
		scheme = "http"
		sub = h.locale
	}

	url := fmt.Sprintf("%s://%s.openfoodfacts.%s%s", scheme, sub, tld, path)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}

	if !h.live {
		request.SetBasicAuth("off", "off")
	}

	return request
}
