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
