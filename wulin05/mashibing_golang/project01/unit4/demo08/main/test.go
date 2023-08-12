package main

import "fmt"

func main() {
	//关于goto的使用: 可以无条件地转移到程序中指定的行,但不建议使用,避免程序流程的混乱
	fmt.Println("hello golang1")
	fmt.Println("hello golang2")
	if 1 == 1 {  //goto一般配合条件结构一起使用
		goto lable1
	}
	fmt.Println("hello golang3")
	fmt.Println("hello golang4")
	fmt.Println("hello golang5")
	lable1:
	fmt.Println("hello golang6")
	fmt.Println("hello golang7")
	fmt.Println("hello golang8")
	fmt.Println("hello golang9")
}
