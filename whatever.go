package goset

import "github.com/ferreiravinicius/goset/hashset"

func main() {
	set := hashset.New()
	set.Add(1)

	set.ToSlice()
}
