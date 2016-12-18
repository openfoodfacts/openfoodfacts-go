package openfoodfacts

import "fmt"

// You can use the GetProduct method to retrieve a Product by barcode.
func ExampleHttpApi_GetProduct() {
	api := NewHttpApiOperator("world", "", "")
	product, err := api.GetProduct("5201051001076")
	if err != nil {
		fmt.Printf("%+v\n", product)
	}
}

// This will enable test mode and connect you to the sandbox server.
func ExampleHttpApi_Sandbox() {
	api := NewHttpApiOperator("world", "", "")
	api.(*HttpApi).Sandbox() // Enable test mode
}

// Create a DataOperator to retrieve and modify database items. HttpApi is a DataOperator that interacts with the
// official HTTP API from openfoodfacts.org.
func Example() {
	NewHttpApiOperator("world", "", "")
}
