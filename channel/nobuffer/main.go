package main

import "fmt"

//func main() {
//	ch := make(chan int)
//
//	ch <- 9
//	fmt.Println("Never print")
//}

func main() {
	ch := make(chan int)

	go func() {
		ch <- 9
	}()

	i := <-ch
	fmt.Println(i)
}
