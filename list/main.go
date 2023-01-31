package main

func main() {

	list := []int{1, 2, 3}
	list2 := []int{5, 6, 7}

	list3 := append(list, list2...)
	println(list3)

}
