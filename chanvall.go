package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main1() {

	syncChan := make(chan struct{})
	go func() {
		for {
			if e, ok := <-mapChan; ok {
				e["count"]++
			} else {
				break
			}

		}
		fmt.Println("stop")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Second)
			fmt.Println("count", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan

}

// 如上所示: mapchan的元素类型属于引用类型, 所以接收方对元素值得副本的修改会影响到发送方持有的原值.
// 每次都应该做一个实验

type Counter struct {
	count int
}

var mapChan2 = make(chan map[string]*Counter, 1)

func main(){
	syncChan := make(chan struct{})
	go func() {
		for {
			if e, ok := <-mapChan2; ok {
				counter := e["count"]
				counter.count ++
			} else {
				break
			}

		}
		fmt.Println("stop")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]*Counter{
			"count": &Counter{},
		}

		for i := 0; i < 5; i++ {
			mapChan2 <- countMap
			time.Sleep(time.Second)
			fmt.Println("count", countMap["count"])
		}
		close(mapChan2)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan

}