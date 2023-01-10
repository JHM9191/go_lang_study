package main

import (
	"fmt"
	"time"
)

func PrintKor() {
	kor := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range kor {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintKor()
	go PrintNumbers()

	//time.Sleep(3 * time.Second)
}
