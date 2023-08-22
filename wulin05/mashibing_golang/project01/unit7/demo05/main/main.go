package main

import (
	"fmt"
)

/*
数组的注意事项:
1. 长度属于类型的一部分
2. 在Go语言中 数组属于值类型,在默认情况下是值传递,因此会进行值拷贝
3. 如想在其他函数中,去修改原来的数组,可以使用引用传递(指针方式)
*/

func main() {
	//定义一个数组:
	// var arr1 = [3]int{3,6,9}
	// fmt.Printf("arr1数组的类型为: %T\n", arr1) //类型是：[3]int

	// var arr2 = [...]int{2,4,6,8}
	// fmt.Printf("arr2数组的类型为: %T\n", arr2) //类型是：[4]int

	var arr3 = [...]int{0: 11, 2: 33, 4: 55, 6: 77, 8: 99}
	fmt.Println(arr3)

	test1(arr3)
	fmt.Println(arr3)
	/*
	    输出还是[11 0 33 0 55 0 77 0 99],没有将arr3[0]的值从11变成111,原因如下:
		test1(arr3)是将arr3数组传递给test01函数,但是同样跟之前的内存分析的道理一样,这是值传递的方式。
		test01函数自己的栈帧会开辟一个arr数组,用以接收arr3数组的值,并且test01函数对arr数组修改值的操作,
		只是test01函数中自己的arr数组的值修改了,等test01函数结束后,就释放空间了,不影响main函数中的arr3这个数组！
	*/
	//那么要实现在其他函数中,去修改原来的数组,可以使用引用传递(指针方式)
	test2(&arr3) //test2函数 传入参数是&arr3,这是arr3的地址,不是arr3的值。
	fmt.Printf("test02函数才真正将arr3[2]的值由原来33修改为333: %v", arr3)
}

func test1(arr [9]int) {
	arr[2] = 333
}

func test2(arr *[9]int) {
	(*arr)[2] = 333
}
