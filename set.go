package gset

type Set interface {
	Add(item interface{}) bool
	Collect() []interface{}
}
