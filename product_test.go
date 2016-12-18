package openfoodfacts_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/openfoodfacts/openfoodfacts-go"
)

func TestProduct(t *testing.T) {
	codes := [...]string{
		"5201051001076",
		"9300650658615",
		"9310155100335",
		"9300601768226",
		"22007377",
		"3222473161867",
		"3242272260059",
		"0737628064502",
		"57626339",
		"5410228196693",
		"5015821151720",
		"5601077161035",
		"3596710352418",
		"8992696419766",
		"3551100749018",
		"5449000179661",
		"3502110000880",
		"8712000031312",
		"6194002510064",
		"3800020430781",
		"5050854517631",
		"5054070608074",
		"8424259826051",
		"5010251168577",
		"4388858946739",
		"4000539770708",
		"3560070976867",
		"2000000039502",
		"0849092103196",
		"7640101710236",
		"3245412470929",
		"3033610048398",
		"0812133010036",
		"8585002476821",
		"3560070805259",
	}

	api := openfoodfacts.NewHttpApiOperator("world", "", "")
	api.(*openfoodfacts.HttpApi).Sandbox()

	for _, code := range codes {
		product, err := api.GetProduct(code)

		if err != nil {
			t.Error("Error during fetching of item", code, err)
		} else if strings.TrimLeft(product.Code, "0") != strings.TrimLeft(code, "0") {
			t.Error("Wrong or different code in retrieved item, expected", code, "got", product.Code, "for item", code, "\n", fmt.Sprintf("%+v", product))
		}
	}
}
