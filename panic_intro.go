package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("exit")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f:", r)
		}
	}()
	//
	a := make([]int, 3)
	fmt.Println(a[200])
	fmt.Println("Returned normally.")
}
