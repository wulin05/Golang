package main

import "fmt"

func main() {
	var num1 int = 10
	fmt.Println(num1)
	var num2 int = (10+22)%3 + 3 - 7 //=右侧的值运算清楚后，再赋值给=的左侧
	fmt.Println(num2)

	var num3 int = 20
	num3 += 10
	fmt.Println(num3)

	//联系:交换两个数的值并输出结果
	var a int = 4
	var b int = 8
	fmt.Printf("a = %v, b = %v \n",a,b)
	//引入一个中间变量
	var t int = a
	a = b
	b = t
	fmt.Printf("a = %v, b = %v",a,b)
}
