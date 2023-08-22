package main

import "fmt"

func main() {
	fmt.Println(add(10, 30))
}

// defer应用场景:
// 主要是defer有延迟执行机制(函数执行最后再执行defer压入栈的语句)
// 比如写了一个连接数据库的函数,函数内定义了defer关闭连接的语句,那么就会在函数执行完后,
// 最后执行defer的语句,完成断开连接数据库的操作,比较省心,省事！
func add(num1 int, num2 int) int {
	//在Golang中,程序遇到defer关键字,不会立即执行defer后的语句,而是将defer后的语句压入一个栈中,然后继续执行函数后面的语句
	defer fmt.Printf("输出num1=%v \n", num1)
	defer fmt.Printf("输出num2=%v \n", num2)
	//栈的特点是：先进后出
	//在函数执行完毕后,从栈中取出语句开始执行,安装先进后出的规则执行语句
	num1 += 100 //第二步实验,添加的这行和下一行的代码, 此时num1: 110
	num2 += 100 //num2: 130,添加这两行代码,是为了说明defer的语句会将代码语句压入栈中,存入的是当前值,不会因为后续的运算,发生变化
	//即, 虽然经过num1 += 100, num2 += 100, 但是defer打印输出的内容还是原先传入的10,30
	var sum int = num1 + num2
	fmt.Printf("输出sum=%v \n", sum)
	return sum
}
