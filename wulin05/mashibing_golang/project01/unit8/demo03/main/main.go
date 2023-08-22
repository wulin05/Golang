package main

import (
	"fmt"
)

/*
切片的遍历
*/

func main() {
	//定义切片：
	// slice :=[]int{11,22,33,44,55,66}
	slice := make([]int, 5, 10)
	slice[0] = 2
	slice[1] = 5
	slice[2] = 7
	slice[3] = 9
	fmt.Println(slice)

	//遍历：
	//普通for循环
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\n", i, slice[i])
	}

	fmt.Println("---------------------------------------")
	//for range方式
	for k, val := range slice {
		fmt.Printf("slice[%v] = %v\n", k, val)
	}

}
