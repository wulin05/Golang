package main

import "fmt"

func main() {
	//可以通过指针改变指向值
	var num int64 = 18
	fmt.Println(num)
	var ptr *int64 = &num
	*ptr = 20
	fmt.Println(num)

	//var ptr *int = num, 指针变量接收的一定是地址值,所以这个报错！
	//var prt *float32 = &num,指针变量的地址不可以不匹配,即&num的类型是*int类型,所以也报错！

	var f1 float32 = 3.1415
	fmt.Println(f1)
	var addr *float32 = &f1
	fmt.Printf("输出指针addr的值是: %v \n", addr)
	fmt.Printf("输出指针addr所指向的内存对应的值是: %v \n", *addr)
	fmt.Printf("指针addr本身在内存中的地址是: %v \n", &addr)

}
