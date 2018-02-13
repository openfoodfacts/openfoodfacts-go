// Copyright © 2016 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

type ProductResults struct {
	Count          int   `json:"count"`
	Products       *[]Product `json:"products"`
}
