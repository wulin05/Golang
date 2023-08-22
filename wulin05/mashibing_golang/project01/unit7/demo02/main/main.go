package main

import (
	"fmt"
)

//数组的引入
//关于下面指针的内容,在unit2的demo14开始有介绍

func main() {
	//定义一个数组: var 数组名 [数组大小]数据类型 , 如下:
	var arr [3]int16

	//获取数组的长度:
	fmt.Println(len(arr))

	//打印数组: int类型的默认值是0
	fmt.Println(arr) // [0 0 0]

	//证明arr中存储的是地址值
	fmt.Printf("arr的地址为: %p\n", &arr)
	//数组的第一个存储空间的地址:
	//从输出结果可以看出: arr和arr[0]的地址是一样的
	fmt.Printf("arr[0]的地址为: %p\n", &arr[0])

	//输出第二个存储空间的地址:
	fmt.Printf("arr[1]的地址为: %p\n", &arr[1])
	//输出第三个存储空间的地址:
	fmt.Printf("arr[2]的地址为: %p\n", &arr[2])
	//从上面arr、arr[0]、arr[1]、arr[2]的地址输出内容可以看出: int16类型是占2个字节

}
