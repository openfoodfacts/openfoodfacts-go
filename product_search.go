// Copyright © 2016 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

type ProductSearch struct {
	SearchTerms string `url:"search_terms"`
	Page int `url:"page"`
	PageSize int `url:"page_size"`
	SortBy string `url:"sort_by"`
}
