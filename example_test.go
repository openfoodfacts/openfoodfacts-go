// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

import (
	"context"
	"fmt"
	"time"
)

// You can use the Client.Product method to retrieve a Product by barcode.
func ExampleClient_Product() {
	api := NewClient("world", "", "")
	product, err := api.Product(context.TODO(), "5201051001076")
	if err != nil {
		fmt.Printf("%+v\n", product)
	}
}

// This will enable test mode and connect you to the sandbox server.
func ExampleClient_Sandbox() {
	api := NewClient("world", "", "")
	api.Sandbox() // Enable test mode
}

// This will set a timeout for the http client
func ExampleClient_Timeout() {
	api := NewClient("world", "", "")
	api.Timeout(10 * time.Second) // Set http client timeout
}

// Create a Client to retrieve and modify database items.
// Client interacts with the official HTTP API from openfoodfacts.org.
func Example() {
	NewClient("world", "", "")
}
