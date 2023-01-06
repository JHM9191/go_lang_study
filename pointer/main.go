package main

import "fmt"

func main() {
	var p1 *int
	var p2 *int

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)

	fmt.Printf("&p1=%v &p2=%v &p1==&p2 => result: %v\n", &p1, &p2, &p1 == &p2)

	a := 1
	p1 = &a
	p2 = &a

	fmt.Printf("p1==p2 => result: %v\n", p1 == p2)

	var b int
	fmt.Println("b=", b)

	c := 3
	var c1 *int = &c
	var c2 *int = c1
	var c3 **int = &c1
	var c4 **int = c3
	var c5 ***int = &c3
	var c6 ***int = c5
	var c7 ****int = &c5
	var c8 ****int = c7
	fmt.Println(c)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
	fmt.Println(c6)
	fmt.Println(c7)
	fmt.Println(c8)
	fmt.Println(*c8)
	fmt.Println(**c8)
	fmt.Println(***c8)
	fmt.Println(****c8)

	d := 7
	var d1 *int = &d
	var d2 **int = &d1

	c5 = &d2
	fmt.Println(d)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(&d2)
	fmt.Println(c5)
	fmt.Println(***c5)
	fmt.Println(c)
}
