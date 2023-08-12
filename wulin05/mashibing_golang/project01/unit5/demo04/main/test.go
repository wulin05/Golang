package main

import "fmt"

//自定义函数: 功能: 交换两个数的值
func exchangeNum (a int, b int) { 
	var t int
	t = a
	a = b
	b = t
}

//这个示例是为了说明在exchangeNum函数交换两个数,并不影响到主函数
//记住, 栈、堆、代码区的概念: 栈一般是存储基本数据类型；堆一般是存储复杂数据类型；代码区是存储用户代码
func main() {
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前的两个数: num1 = %v, num2 = %v \n", num1, num2)
	exchangeNum(num1, num2)
	fmt.Printf("交换后的两个数: num1 = %v, num2 = %v \n", num1, num2)

}
