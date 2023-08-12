package main

import "fmt"

func main() {
	//布尔类型也叫bool类型，bool类型数据只允许取值true和false
	var a int
	var b float32
	var c float64
	var d bool
	var e string
	fmt.Printf("输出int类型的默认值: %d\n", a)
	fmt.Printf("输出float32类型的默认值: %f\n", b)
	fmt.Printf("输出float64类型的默认值: %f\n", c)
	fmt.Printf("输出int类型的默认值: %t\n", d)
	fmt.Printf("输出int类型的默认值: %q", e)
}
