package main

import "fmt"

func main() {
	var age = 18
	//&+变量名 就可以获取这个变量的内存地址
	fmt.Println(&age)

	//定义一个指针变量:
	//ptr指针变量，*int代表一个指针类型(可以理解为 指向int类型的指针)
	var ptr *int = &age
	fmt.Println("ptr的值就是age这个变量在内存的地址: ", ptr)
	fmt.Println("ptr这个指针变量的本身内存地址: ", &ptr)
	fmt.Printf("获取ptr这个指针或者这个地址指向的那个变量(age)的数据: %v \n", *ptr)
	//所以可以通过prt的指针地址,来对age变量进行修改
	*ptr = 20
	fmt.Printf("现在age的值变成了: %v", age)
}
