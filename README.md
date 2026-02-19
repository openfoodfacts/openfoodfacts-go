# openfoodfacts-go

[![Build Status](https://travis-ci.org/openfoodfacts/openfoodfacts-go.svg?branch=master)](https://travis-ci.org/openfoodfacts/openfoodfacts-go)

## About

This library is for accessing [Open Food Facts](https://world.openfoodfacts.org/) data for food product, ingredients and nutritional data from within your go application. You can also get the Nutri-Score and the Green-Score

Use of the Open Food Facts database is subject to the [Open Food Facts Terms of Use](https://world.openfoodfacts.org/terms-of-use), please read them before using this library in your application.

This library is copyright © 2019-2025 OpenFoodFacts. All rights reserved. Use of this library is governed by the MIT license which can be found in the [LICENSE.txt](LICENSE.txt) file.

## Documentation

[Package documentation](https://godoc.org/github.com/openfoodfacts/openfoodfacts-go)

[OpenFoodFacts API details](https://wiki.openfoodfacts.org/Project:API)

[Go API wiki page](https://wiki.openfoodfacts.org/API/Go)

## Usage details

This is a very simple example how to use the library in you go program.

~~~go
package main

import (
	"github.com/openfoodfacts/openfoodfacts-go"
)

func main() {
	api := openfoodfacts.NewClient()
	product, err := api.Product("0737628064502")
}

~~~

There are runnable examples in the `examples` subdirectory.

## Authors
Ken Allan

## Third party applications
* Feel [free to open a PR to add your application in this list](https://github.com/openfoodfacts/openfoodfacts-go/edit/develop/REUSERS.md).
* Please get in touch at reuse@openfoodfacts.org
* We are very interested in learning what the Open Food Facts data is used for. It is not mandatory, but we would very much appreciate it if you tell us about your re-uses (https://forms.gle/hwaeqBfs8ywwhbTg8) so that we can share them with the Open Food Facts community. You can also fill this form to get a chance to get your app featured: https://forms.gle/hwaeqBfs8ywwhbTg8
* Make sure you comply with the OdBL licence, mentioning the Source of your data, and ensuring to avoid combining non free data you can't release legally as open data. Another requirement is contributing back any product you add using this SDK.

