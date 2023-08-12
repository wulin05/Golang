package main

import "fmt"

func main() {
	//布尔类型也叫bool类型，bool类型数据只允许取值true和false
	//布尔类型占一个字节
	//布尔类型适用于逻辑运算，一般用于程序的流程控制
	//测试布尔类型的数值
	var flag01 bool = true
	fmt.Println(flag01)

	var flag02 bool = false
	fmt.Println(flag02)

	var flag03 bool = 5 < 9
	fmt.Println(flag03)

}
