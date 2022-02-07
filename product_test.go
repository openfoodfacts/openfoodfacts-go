// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/openfoodfacts/openfoodfacts-go"
)

func TestProduct_Unmarshalling(t *testing.T) {
	cases := []struct {
		fixture   string
		returnErr bool
		name      string
	}{
		{"testdata/product/5201051001076.json", false, "Unmarshalling_5201051001076"},
		{"testdata/product/9300601768226.json", false, "Unmarshalling_9300601768226"},
		{"testdata/product/9300650658615.json", false, "Unmarshalling_9300650658615"},
		{"testdata/product/9310155100335.json", false, "Unmarshalling_9310155100335"},
		{"testdata/product/22007377.json", false, "Unmarshalling_22007377"},
		{"testdata/product/3222473161867.json", false, "Unmarshalling_3222473161867"},
		{"testdata/product/3242272260059.json", false, "Unmarshalling_3242272260059"},
		{"testdata/product/0737628064502.json", false, "Unmarshalling_0737628064502"},
		{"testdata/product/5410228196693.json", false, "Unmarshalling_5410228196693"},
		{"testdata/product/5015821151720.json", false, "Unmarshalling_5015821151720"},
		{"testdata/product/5601077161035.json", false, "Unmarshalling_5601077161035"},
		{"testdata/product/3596710352418.json", false, "Unmarshalling_3596710352418"},
		{"testdata/product/8992696419766.json", false, "Unmarshalling_8992696419766"},
		{"testdata/product/3551100749018.json", false, "Unmarshalling_3551100749018"},
		{"testdata/product/5449000179661.json", false, "Unmarshalling_5449000179661"},
		{"testdata/product/3502110000880.json", false, "Unmarshalling_3502110000880"},
		{"testdata/product/8712000031312.json", false, "Unmarshalling_8712000031312"},
		{"testdata/product/6194002510064.json", false, "Unmarshalling_6194002510064"},
		{"testdata/product/3800020430781.json", false, "Unmarshalling_3800020430781"},
		{"testdata/product/5050854517631.json", false, "Unmarshalling_5050854517631"},
		{"testdata/product/5054070608074.json", false, "Unmarshalling_5054070608074"},
		{"testdata/product/8424259826051.json", false, "Unmarshalling_8424259826051"},
		{"testdata/product/5010251168577.json", false, "Unmarshalling_5010251168577"},
		{"testdata/product/4388858946739.json", false, "Unmarshalling_4388858946739"},
		{"testdata/product/4000539770708.json", false, "Unmarshalling_4000539770708"},
		{"testdata/product/3560070976867.json", false, "Unmarshalling_3560070976867"},
		{"testdata/product/0849092103196.json", false, "Unmarshalling_0849092103196"},
		{"testdata/product/7640101710236.json", false, "Unmarshalling_7640101710236"},
		{"testdata/product/3245412470929.json", false, "Unmarshalling_3245412470929"},
		{"testdata/product/3033610048398.json", false, "Unmarshalling_3033610048398"},
		{"testdata/product/0812133010036.json", false, "Unmarshalling_0812133010036"},
		{"testdata/product/8585002476821.json", false, "Unmarshalling_8585002476821"},
		{"testdata/product/3560070805259.json", false, "Unmarshalling_3560070805259"},
		{"testdata/product/7311041026670.json", false, "Unmarshalling_7311041026670"},
		{"testdata/product/ingredients_n_as_int.json", false, "Unmarshalling_IntIngredientNAsInt"},
		{"testdata/product/ingredients_n_as_string.json", false, "Unmarshalling_IntIngredientNAsString"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			productFile, err := os.Open(tc.fixture)
			if err != nil {
				t.Error("Error when opening test file", tc.fixture)
				return
			}
			defer productFile.Close()

			byteProducts, err := ioutil.ReadAll(productFile)
			if err != nil {
				t.Error("Error when reading test data", tc.fixture)
				return
			}

			var pr openfoodfacts.ProductResult
			err = json.Unmarshal(byteProducts, &pr)
			if err != nil {
				t.Error("Error when unmarshalling test data", tc.fixture)
				return
			}
		})
	}
}
