package main

import "fmt"

type Rabbit struct {
	width  int
	height int
}

func (r Rabbit) info() int {
	return r.width * r.height
}

func main() {
	rabbit := &Rabbit{5, 6}
	size := rabbit.info()
	fmt.Println(size)
}
