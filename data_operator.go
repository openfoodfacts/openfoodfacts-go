package openfoodfacts

// A DataOperator is capable of retrieving details of food items from a database and modifying them.
// Typically this will be either from the official OpenFoodFacts API, a local database, or may be some form of caching
// layer above one of these other methods.
//
// HttpApi is an example of a DataOperator that uses the OpenFoodFacts web API.
//
// GetProduct will return a pointer to a Product given a item code.
type DataOperator interface {
	GetProduct(code string) (*Product, error)
}
