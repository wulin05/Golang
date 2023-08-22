package main

import (
	"fmt"
)

/*
6. 切片的拷贝
*/

func main() {
	//新建四个切片a, b, c, d
	a := []int{1, 4, 7, 3, 6, 9}
	var b []int = []int{11, 21, 31, 41, 51, 61}
	c := make([]int, 5)
	d := make([]int, 10)

	//拷贝：将a里面的数据赋值给c
	copy(c, a)     //a,c切片的底层都是数组,那么其实就是将a对应数组的元素复制到b中对应的数组中
	fmt.Println(c) // [1 4 7 3 6],c没有全部接收a的数据
	//拷贝：将b里面的数据分支给d
	copy(d, b)
	fmt.Println(d) // [11 21 31 41 51 61 0 0 0 0],d全部接收b的数据

}
