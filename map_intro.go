package main

import "fmt"

var m2 map[string]string // 此时他是一个nil

func main() {
	m2 = make(map[string]string) // 创建一个非nil的ma
	m2["a"] = "2"
	m2["b"] = "3"
	fmt.Println(m2)
}
