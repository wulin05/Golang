package main

import "fmt"

//既然上一节说的基本数据类型和数组的值传递方式无法改变函数(test函数无法改变main函数里的值)外的值
//那么要想能改变的话,就要使用指针的方式(通过引用指针,改变指针指向内存的数据)
func test (num *int) {
	//从main函数传入的是m的地址,那么test函数的创建num的栈帧存的是传入的m的地址
	//这时候num的值是m的地址的话,如何通过这个地址能够修改m的值呢？
	*num = 55    //通过num存的m的地址,找到m,然后修改m的值为55！
	fmt.Println("test函数中num存的是main函数m的地址: ", num)
	fmt.Println("注意输出中,第一行是main函数m的地址,与第二行的test函数输出num的值是一样的!")
	fmt.Println("test函数中num自己本身的内存地址: ", &num)
}

func main() {
	var m int = 10
	fmt.Println(&m) //输出变量m的地址值
	test(&m) //传入的是变量m的地址,test函数通过这个地址*m,就能改变main函数的这个m变量的存储值
	fmt.Printf("原来main中的m的值经过test函数的处理后,m的值变为: %v", m) //改变后,m的输出结果就变成了修改后的值了!
}

