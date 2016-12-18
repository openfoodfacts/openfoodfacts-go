# openfoodfacts-go

[![Build Status](https://travis-ci.org/openfoodfacts/openfoodfacts-go.svg?branch=master)](https://travis-ci.org/openfoodfacts/openfoodfacts-go)

## About

This library is for accessing [OpenFoodFacts](http://world.openfoodfacts.org/) data for food product, ingredients and nutritional data from within your go application.

Use of the OpenFoodFacts database is subject to the [OpenFoodFacts Terms of Use](http://world.openfoodfacts.org/terms-of-use), please read them before using this library in your application.


## Documentation

[Package documentation](https://godoc.org/github.com/openfoodfacts/openfoodfacts-go)

[OpenFoodFacts API details](http://en.wiki.openfoodfacts.org/Project:API)

[Go API wiki page](http://en.wiki.openfoodfacts.org/API/Go)

## Usage details

This is a very simple example how to use the library in you go program.

~~~go
package main

import (
	"github.com/openfoodfacts/openfoodfacts-go"
)

func main() {
	api := openfoodfacts.NewHttpApiOperator("world", "", "")
	product, err := api.GetProduct("0737628064502")
}

~~~

There are runnable examples in the `examples` subdirectory.
