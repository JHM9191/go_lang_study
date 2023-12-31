package main

import (
	"fmt"
)

type Tree[T interface{}] struct {
	left, right *Tree[T]
	data        T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	//...
}

var t Tree[string]

//심지어 두 타입이 연계될 수도 있다.
func Scale[S ~[]E, E constraints.Integer](s S, sc E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func min[T](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	x := min[int](3, 2)
	fmt.Println(x)
}
