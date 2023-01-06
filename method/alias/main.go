package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10
	fmt.Println(a.add(20))

	b := 50
	fmt.Println(myInt(b).add(50))
}
