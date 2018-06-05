package main

import "fmt"

func main() {

	var values []int        // 声明 但是不指定长度， 这个时候slice长度为0，还不能直接放数据
	values = make([]int, 5) //调用内置函数make分配存储空间

	for i := range values {
		values[i] = i + 1
	}

	fmt.Println(values)

	values = make([]int, 6, 10) // 创建长度为10的数组， 仅仅返回前五
	fmt.Println(values)

	values = append(values, 7, 8)
	values = append(values, 7, 8)
	values = append(values, 7, 8)
	values = append(values, 7, 8)
	values = append(values, 7, 8)
	fmt.Println(values)

	// copy  如果目标的长度小于源的长度， 那么不会拷贝所有源

	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	slice3 := make([]int, 3)
	slice4 := make([]int, 4)

	copy(slice2, slice1)
	copy(slice3, slice1)
	copy(slice4, slice1)
	fmt.Println(slice2, "slice2")
	fmt.Println(slice3, "slice3")
	fmt.Println(slice4, "slice4")

	}
