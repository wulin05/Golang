package main

import (
	"fmt"
)

//二维数组的遍历

func main() {
	//如果是这样定义的话：
	var arr0 = [][4]int{{0}, {1, 2}, {3, 4, 5}, {6, 7, 8, 9}}
	fmt.Println(arr0)

	fmt.Println("---------------------------------------------")
	var arr = [][]int{{}, {1, 2}, {3, 4, 5}, {6, 7, 8, 9}}
	fmt.Println(arr)
	//下面4个打印是我自己想知道的,视频没有说明这个！
	fmt.Printf("arr的类型: %T, 数组arr长度是: %v\n", arr, len(arr))
	fmt.Printf("arr[0]的类型是: %T, 数组arr[0]长度是: %v\n", arr[0], len(arr[0]))
	fmt.Printf("arr[1]的类型是: %T, 数组arr[0]长度是: %v\n", arr[1], len(arr[1]))
	fmt.Printf("arr[2]的类型是: %T, 数组arr[2]长度是: %v\n", arr[2], len(arr[2]))
	fmt.Printf("arr[3]的类型是: %T, 数组arr[3]长度是: %v\n", arr[3], len(arr[3]))

	fmt.Println("---------------------------------------------")
	//普通for循环
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
			//注意 这里是用fmt.Print,不是之前的fmt.Println和fmt.Printf了!
		}
		fmt.Println()
	}

	fmt.Println("---------------------------------------------")
	//键值循环(for range),循环的本质是：arr[key][k]=v
	for key, val := range arr {
		for k, v := range val {
			fmt.Printf("arr[%v][%v] = %v\t", key, k, v)
		}
		fmt.Println()
	}
}
