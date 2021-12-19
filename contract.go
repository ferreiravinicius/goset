package gset

type Set interface {
	Add(item interface{}) bool
	Slice() []interface{}
}
