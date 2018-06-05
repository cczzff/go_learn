package main

import "fmt"

var v = "1, 2, 3"

func main() {
	v := []int{1, 2, 3}
	if v != nil {
		var v = 123
		fmt.Println(v)
	}
	fmt.Println(v)
	main2()
}

func main2() {
	fmt.Println(v)
}
