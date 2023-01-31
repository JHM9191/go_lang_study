package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)
	wg.Wait()
	fmt.Println("Close Factory")
}

func MakeBody(tireCh chan *Car) { // 차체 생상
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		fmt.Println("MakeBody")
		select {
		case <-tick:
			// Make a body
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after: // 10초 뒤 종료
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) { // 바퀴 설치
	for car := range tireCh {
		fmt.Println("InstallTire")
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) { // 도색
	for car := range paintCh {
		fmt.Println("PaintCar")
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
