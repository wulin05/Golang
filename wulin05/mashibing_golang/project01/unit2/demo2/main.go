package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var num1 = 28
	fmt.Printf("num1的类型是: %T", num1)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num1)) //unsafe.Sizeof是用于求出num1占用的字节数

	//表示学生的年龄
	var age uint8 = 18
	//或者 var age byte = 18
	fmt.Println(age)
}
