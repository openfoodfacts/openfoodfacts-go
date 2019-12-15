// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

type NutrientLevel struct {
	Sugars       string `json:"sugars"`
	SaturatedFat string `json:"saturated-fat"`
	Fat          string `json:"fat"`
	Salt         string `json:"salt"`
}
