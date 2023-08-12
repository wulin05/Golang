package main

import "fmt"

/*
函数的介绍:
1.提供代码的复用性,减少代码的冗余,代码的维护性也提高了
2.函数的定义是: 为完成某一功能的程序指令(语句)的集合,成为函数
3.基本语法:
func 函数名(形参列表)(返回值类型列表) {
	执行语句...
	return + 返回值列表
}

*/

//自定义函数: 功能: 两个数相加
func cal (num1 int, num2 int) (int) { //如果返回值类型就一个的话,那么(int)的括号是可以省略的
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}

func main() {
	//功能: 10 + 20
	// var num1 int = 10
	// var num2 int = 20
	// sum := num1 + num2
	// fmt.Println(sum)
	// fmt.Printf("sum的类型是: %T, sum = %v", sum, sum)

	//调用函数cal
	sum1 := cal(10, 20)
	fmt.Println(sum1)

	//例子二:
	var num3 int = 30
	var num4 int = 50
	sum2 := cal(num3, num4)
	fmt.Println(sum2)
}
