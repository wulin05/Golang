package main

import "fmt"

func main() {
	//定义一个变量
	var age int = 38
	fmt.Println("age对应的存储空间的地址是:",&age)

	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr这个指针指向的具体数值为:",*ptr)

}
