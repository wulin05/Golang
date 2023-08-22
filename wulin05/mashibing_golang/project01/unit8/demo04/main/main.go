package main

import (
	"fmt"
)

/*
切片的注意事项：
1. 切片定义后不可以直接使用，需要让其引用到一个数组，或者make一个空间供切片使用
2. 切片不能越界
3. 切片的简写方式：
   var slice = arr[0:end]  =====> var slice = arr[:end]
   var slice = arr[start:len[arr]]  =======> var slice = arr[start:]
   var slice = arr[0:len[arr]]   ======> var slice = arr[:]

4. 切片可以继续切片
*/

func main() {
	//定义切片：
	var slice []int
	fmt.Println(slice) //[]

	//make方式出来的切片
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1) //[0 0 0 0 0]

	//定义一个数组
	var intarr [6]int = [6]int{1, 3, 5, 7, 9}
	var slice2 []int = intarr[1:4]
	fmt.Println(slice2) // [3 5 7]

	//越界测试：panic: runtime error: index out of range [3] with length 3
	// fmt.Println(slice2[3])

	//切片的简写方式：
	slice3 := intarr[1:]
	fmt.Println(slice3) // [3 5 7 9 0]

	//切片再切片
	slice4 := slice3[1:3]
	fmt.Println(slice4) // [5 7]

	//对slice4修改值，会影响intarr、slice2、slice3么？
	//答案是肯定的
	slice4[1] = 777
	fmt.Println(slice4) // [5 777]
	fmt.Println(slice3) // [3 5 777 9 0]
	fmt.Println(slice2) // [3 5 777]
	fmt.Println(intarr) // [1 3 5 777 9 0]
}
