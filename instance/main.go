package main

import "fmt"

type Data struct {
}

func main() {
	var data Data
	var p *Data = &data

	fmt.Println(p)

	var s *Data = &Data{}
	fmt.Println(s)

	var t = &Data{}
	fmt.Println(t)

	u := &Data{}
	fmt.Println(u)

	var data1 Data
	var data2 Data = data1
	var data3 Data = data2
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)
}
