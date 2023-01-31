package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
	//fmt.Println(<-sendch) // Error
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
}

//10
