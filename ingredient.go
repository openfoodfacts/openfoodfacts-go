// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

type Ingredient struct {
	ID         string `json:"id"`
	Rank       int    `json:"rank,omitempty"`
	Text       string `json:"text"`
	Vegan      string `json:"vegan,omitempty"`
	Vegetarian string `json:"vegetarian,omitempty"`
	Percent    string `json:"percent,omitempty"`
}
