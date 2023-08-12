package main

import "fmt"

func main() {
	//与逻辑: &&,也叫短路与:
	//只要第一个数值/表达式的结果是false,那么后面的表达式等就不用运算了,结果就是false -->提高运算效率
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	//或逻辑: ||
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	//非逻辑: ! 取想法的结果
	fmt.Println(!true)
	fmt.Println(!false)
}
