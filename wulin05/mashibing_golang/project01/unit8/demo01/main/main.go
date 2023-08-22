package main

import (
	"fmt"
)

/*
切片的引入：切片其实是一个结构体,包含三部分：底层数组的指针、切片的长度、切片的容量
切片是构建在数组之上,数组有特定的用处,但是却有些呆板(数组长度固定不可变),所以在Go语言代码数组不常见
切片是对数组一个连续片段的引用,所以切片是一个引用类型。
切片提供了一个相关数组的动态窗口。

定义一个切片名字为slice,[]动态变化的数组,长度不写,int类型,intarr是原数组
//[1:3]切片 - 切出数组的一段片段  索引: 从1开始,到3结束（不包含3）
*/

func main() {
	//定义一个数组,关于数组的4中不同初始化方式,可以查看unit7的demo04章节
	var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	fmt.Println(intarr)

	// intarr1 := [...]int{2,5,8}
	// fmt.Println(intarr1)

	var slice []int = intarr[1:3] //记住：slice []int, 切片的[]是不写长度的,否则不就变成数组了
	fmt.Println(slice)

	// slice1 := intarr[1:3]
	// fmt.Println(slice1)

	//切片元素个数:
	fmt.Println("slice元素的个数:", len(slice))

	//获取切片的容量,容量可以动态变化
	fmt.Println("slice的容量:", cap(slice))

	//下面测试切片slice是一个结构体，其中结构体中有一个是存放底层数组指针的信息
	//那么slice存放底层数组指针信息应该就是intarr[1]的地址！
	fmt.Printf("intarr[1]的内存地址是: %p\n", &intarr[1])
	fmt.Printf("slice[0]的内存地址是: %p\n", &slice[0])
	//可以明显地看出内存地址相同

}
