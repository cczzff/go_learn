package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("map:%v\n", m)
	info = fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
}

var m = map[int]string{1: "a", 2: "b", 3: "c"}

var info string

func main(){
	fmt.Println(info)
}

