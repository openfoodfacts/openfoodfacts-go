// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

// This is a go library used to access the OpenFoodFacts.org database for food product, ingredients and
// nutritional data from within your go application.
//
// The main method of using this library is to create a DataOperator and call methods on it.
package openfoodfacts

import "errors"

// A DataOperator is capable of retrieving details of food items from a database and modifying them.
// Typically this will be either from the official OpenFoodFacts API, a local database, or may be some form of caching
// layer above one of these other methods.
//
// Client is an example of a DataOperator that uses the OpenFoodFacts web API.
type DataOperator interface {
	// GetProduct will return a pointer to a Product, given a item code.
	GetProduct(code string) (*Product, error)
}

var (
	// ProductRetrievalError is an error returned by DataOperator.GetProduct when the product could not be
	// retrieved successfully, it is not a transient error, the product does not exist.
	ProductRetrievalError = errors.New("Product retrieval failure")

	// AnonymousUserError is an error returned by DataOperator methods that require a valid user account, but none
	// was provided when the operator was instantiated.
	AnonymousUserError = errors.New("Action requires user account")
)
