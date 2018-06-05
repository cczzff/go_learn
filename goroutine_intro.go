package main

import (
	"fmt"
	"time"
)

func main() {
	// 最终输出的是ali
	name := "xiaolongxia "
	go func() {
		fmt.Println(name, "gogogo")

	}()
	name = "ali"
	time.Sleep(time.Millisecond)

	// 最终输出的是ccc
	names := []string{"a", "b", "c"}
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
	}

	time.Sleep(time.Millisecond)

	// 这样就可以执行 a b c 了
	names = []string{"a", "b", "c"}
	for _, name := range names {
		go func(who string) {
			fmt.Println(who, "who")
		}(name)

	}
	time.Sleep(time.Millisecond)


}
