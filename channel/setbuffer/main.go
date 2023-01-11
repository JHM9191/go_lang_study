package main

import "fmt"

func main() {
	var messages chan string = make(chan string, 2)
	messages <- "A"
	messages <- "B"
	//messages <- "C"
	fmt.Println("Print me")
}
