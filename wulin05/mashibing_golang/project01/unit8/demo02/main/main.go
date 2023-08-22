package main

import (
	"fmt"
)

/*
切片的定义：
上一节已经有两个方式定义切片
本章节使用make方式定义切片： make(Type,size integer Type) Type
基本语法：var 切片名[type=make([],len,[cap])]
*/

func main() {
	//定义一个切片make方式：
	//make方式是底层先创建了一个数组,对外不可见,所以不可以直接操作这个数组,
	//要通过slice去间接访问各个元素,即不可以直接对数组进行维护/操作
	slice := make([]int, 4, 20)
	fmt.Println(slice)                // [0 0 0 0]
	fmt.Println("切片的长度:", len(slice)) //4
	fmt.Println("切片的容量:", cap(slice)) //20

	//赋值操作：
	slice[0] = 66
	slice[1] = 77
	fmt.Println(slice)

	fmt.Println("-----------------------------------------------")
	//方式二：定义一个切片,直接就指定具体数组,使用原理类似make的方式
	//也是底层创建了一个数组,但是对外不可见,只能通过切片来维护/操作
	slice2 := []int{1, 4, 7}
	fmt.Println(slice2)
	fmt.Println("切片的长度:", len(slice2)) //3
	fmt.Println("切片的容量:", cap(slice2)) //3

}
