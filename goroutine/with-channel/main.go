package main

import (
	"fmt"
	"time"
)

func PrintKor(ch chan int) {
	kor := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range kor {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
	ch <- 1
}

func PrintNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	ch <- 2
}

type Data struct {
	i int
}

func PrintMe(data *Data, ch chan int) {
	data.i = 1
}
func main() {
	func() {
		data := Data{}
		c1 := make(chan int)
		go PrintMe(&data, c1)
		cc1 := <-c1
		fmt.Println(cc1)
		c2 := make(chan int)
		go PrintMe(&data, c2)
		cc2 := <-c2
		fmt.Println(cc2)
	}()
	//time.Sleep(3 * time.Second)
}
