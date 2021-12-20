package goset

// A collection that contains no duplicate elements.
type Set interface {

	// Adds the item to this set if it is not already present.
	// Returns true if added and false otherwise.
	Add(item interface{}) bool

	// Creates a slice containing all items.
	Collect() []interface{}

	// Returns true if provided item exists in this set.
	Contains(item interface{}) bool

	// Removes the provided item from this set if it exists.
	// Returns true if removed and false otherwise.
	Remove(item interface{}) bool

	// Returns true if all provided items exists in this set.
	ContainsAll(items ...interface{}) bool

	// Returns the length of this set.
	Len() int
}
