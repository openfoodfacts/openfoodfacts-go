package openfoodfacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpApi struct {
	locale   string
	username string
	password string
	live     bool
}

// NewHttpApiRetriever returns a DataOperator that is capable of talking to the official OpenFoodFacts database via
// the HTTP API, or the dev server if live is false.
//
// locale should be one of "world" or the country level code for the locale you wish to use.
// username and password should be set to your OpenFoodFacts credentials if you need WRITE access, else "".
//
// Call .Sandbox() on the returned operator if you are testing functionality via the sandbox environment.
func NewHttpApiOperator(locale, username, password string) DataOperator {
	return &HttpApi{
		locale:   locale,
		username: username,
		password: password,
		live:     true,
	}
}

// A ProductRetrievalError is returned by FetchProduct when the resulting status is not 1.
var ProductRetrievalError = errors.New("Product retrieval failure")

// GetProduct returns a new Product for the given code, retrieved from the server.
//
// It will return an error on a failed retrieval, if the retrieval is successful but the status is not 1, then will
// return ProductRetrievalError.
func (h *HttpApi) GetProduct(code string) (*Product, error) {
	request := h.newRequest("GET", "/api/v0/product/%s.json", code)

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

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

// Configures this operator to use the sandbox server at world.openfoodfacts.net instead of the live server.
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
