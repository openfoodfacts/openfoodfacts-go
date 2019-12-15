// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

type ProductResult struct {
	StatusVerbose string   `json:"status_verbose"`
	Status        int      `json:"status"`
	Code          string   `json:"code"`
	Product       *Product `json:"product"`
}
