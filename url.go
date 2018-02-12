// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

import (
	"fmt"
	"net/url"
	"strings"
)

// URL allows URL to be properly un/marshaled from a JSON string.
type URL struct {
	url.URL
}

func (u *URL) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	parsed, err := url.Parse(s)
	if err != nil {
		return err
	}

	u.URL = *parsed
	return nil
}

func (u *URL) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", u.URL.String())), nil
}
