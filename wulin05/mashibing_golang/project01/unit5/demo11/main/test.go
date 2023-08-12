package main

import (
	"fmt"
)

/*
	使用包的原因:
	1.我们不可能把所有的函数放在同一个源文件中,可以分门别类的把函数放在不同的源文件中
	2.解决同名问题: 两个人都想定义一个同名的函数,在一个文件中是不可以定义相同名字的函数的。此时可以用包来区分
	  //貌似函数重载,go不支持函数重载,也就是在一个包里有两个同名的函数....
*/


func test01 (num1 int, num2 int) (int, int) {
	result01 := num1 + num2
	result02 := num1 - num2
	return result01,result02
}

func test02 (num1 int, num2 int) (sub int, sum int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	a, b := test01(1, 2)
	fmt.Println(a,b)
	c, d := test02(10, 20)
	fmt.Println(c,d)
}

