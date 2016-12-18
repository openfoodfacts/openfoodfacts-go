// A simple application to demonstrate looking up a product from the OpenFoodFacts database using the EAN number
// from a barcode.
//
// To build this code, simply:
//   go get
//   go build
//
// To specify the barcode to use, run the program with the "-code" flag, eg:
//   ./simple-get-product -code "0737628064502"
//
// For more detail:
//   ./simple-get-product -help
package main

import (
	"flag"
	"fmt"

	"github.com/openfoodfacts/openfoodfacts-go"
)

var (
	code string
)

func init() {
	flag.StringVar(&code, "code", "0737628064502", "Supply the barcode of the item you wish to retrieve")
	flag.Parse()
}

func main() {
	api := openfoodfacts.NewHttpApiOperator("world", "", "")
	product, err := api.GetProduct(code)
	if err == nil {
		fmt.Printf("%+v\n", product)
	} else {
		fmt.Println("Error:", err)
	}
}
