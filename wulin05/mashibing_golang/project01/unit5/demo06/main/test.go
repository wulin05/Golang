package main

import "fmt"

//定义一个函数,函数参数为: 可变参数... 参数的数量可变
//args...int 可以传入任意多个数量的int类型的数据
func test (args...int) {
	//在函数内部处理可变参数的时候,将可变参数当做切片来处理
	//遍历可变参数
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func main() {
	test()
	fmt.Println("------------------")
	test(3)
	fmt.Println("------------------")
	test(21, 121, 333, 111, 99)
}
