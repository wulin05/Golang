package main

import (
	"fmt"
)

/*
	该代码是为了说明以下内容: 
	传统写法要求test01: 返回值和返回值的类型对应,顺序不能乱
	新的写法test02: 支持对函数返回值命名,那里面顺序就无所谓了,顺序不用对应
*/

//
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
	var a int
	var b int
	var c int
	var d int
	a, b = test01(50, 30)
	c, d = test02(500, 300)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

