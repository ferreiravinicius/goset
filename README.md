Implementation of set for Golang.  
Set is a collection that contains no duplicate elements.  

Curently in **early development** phase.

## Install

```bash
$ go get -u github.com/ferreiravinicius/goset
```

## Hash Set

*Creating a new set*
```golang
import "github.com/ferreiravinicius/goset/hashset"

set := hashset.New() // Set{}
set := hashset.From(1, 2, 3) // Set{1, 2, 3}
set := hashset.New(100) // creates new map with size of 100
```

*Iterating over a set using `ForEach`*
```golang
set := hashset.From(1, 2, 3) 
set.ForEach(func(item interface{}) { 
  n := item.(int) 
  // doSomething(n)
})
```




