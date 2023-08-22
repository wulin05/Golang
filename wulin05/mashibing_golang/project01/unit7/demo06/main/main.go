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
	//定义一个二维数组
	var arr [2][3]int16
	fmt.Println(arr) // 结果：[[0 0 0] [0 0 0]]
	/*
		想象一下,其实二维数组的结构和一维数组结构是一样的:
		一维数组arr01[2]：[0,0,0]
		二维数组arr02[2][2]：[[0,0,0],[0,0,0],[0,0,0]]
		从内存上分析：
		一维数组arr01的地址就是 &arr01,也是arr01[0]
		二维数组arr02的地址就是 &arr02,也是arr02[0],也是arr02[0][0]
	*/
	//下面关于arr的地址：&arr == &arr[0] == &arr[0][0]
	fmt.Printf("arr的地址是: %p\n", &arr)
	fmt.Printf("arr[0]的地址是: %p\n", &arr[0])
	fmt.Printf("arr[0][0]的地址是: %p\n", &arr[0][0])

	//同理arr[1][0]的地址: &arr[1] == &arr[1][0]
	fmt.Printf("arr[1]的地址是: %p\n", &arr[1])
	fmt.Printf("arr[1][0]的地址是: %p\n", &arr[1][0])
	//再注意看结果,可以发现int16是2个字节,所以 &arr[1] == &arr[0] + 6
	//因为arr[2]里的每个数据又是一个数组,有三个元素：arr[2][3],所以从&arr[0]到&arr[1],地址多了6个字节

	//从上面的一系列分析,是让你明白,其实不管是几维数组,最终都可以看成是一维数组！！！
	//赋值：
	arr[0][1] = 1
	arr[0][2] = 2
	arr[1][1] = 1
	arr[1][2] = 2
	fmt.Println(arr) // [[0 1 2] [0 1 2]]

	//二维数组的初始化操作：就是一维的变形
	var arr1 [2][3]int = [2][3]int{{11, 22, 33}, {44, 55, 66}}
	fmt.Println(arr1)

	var arr2 = [3][2]int{{0, 1}, {2, 3}, {4, 5}}
	fmt.Println(arr2)

	var arr3 = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//或者：var arr3 = [...][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(arr3)

	//注意下面两种定义后,不同之处:
	var arr4 = [][]int{{0: 100}, {1: 200}, {2: 300}, {3: 400}, {4: 500}}
	fmt.Println(arr4) // [[100] [0 200] [0 0 300] [0 0 0 400] [0 0 0 0 500]]
	var arr5 = [][5]int{{0: 100}, {1: 200}, {2, 300}, {3: 400}, {4: 500}}
	fmt.Println(arr5)
}