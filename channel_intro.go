package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() { // 用来演示接收工作
		<-syncChan1
		fmt.Println("接收到1个信号, 停留一秒")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("接收到 string", elem)
			} else {
				fmt.Println("关闭")
				break
			}

		}
		syncChan2 <- struct{}{}
	}()

	go func() { // 用于演示发送
		for _, e := range []string{"a", "b", "c", "d"} {
			strChan <- e
			fmt.Println("send", e)
			if e == "c" {
				fmt.Println("signal")
				syncChan1 <- struct{}{}
			}
			time.Sleep(time.Second * 2)
		}
		close(strChan)
		syncChan2 <- struct{}{}

	}()

	<-syncChan2
	<-syncChan2
}
