package goset

// A collection that contains no duplicate elements.
// Offers constant time performance for the basic operations.
type Set interface {

	// Adds the item to this set if it is not already present.
	// Returns true if added and false otherwise.
	Add(item interface{}) bool

	// Creates a slice containing all items.
	Collect() []interface{}
}
