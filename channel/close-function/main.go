package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i // chnl 에 값을 안 넣으면 ok가 false
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	value, ok := <-ch // 성공: ok가 true | 실패: ok가 false
	fmt.Println(value, ok)

	for v := range ch {
		fmt.Println("Received ", v)
	}
}

// https://hoony-gunputer.tistory.com/entry/goroutine-and-channel
