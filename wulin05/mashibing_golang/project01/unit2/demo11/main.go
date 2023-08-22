package main

import "fmt"

//基本数据类型转换成字符串类型
func main() {

	//n1 := 19
	var n1 int = 19
	var n2 float32 = 4.78
	var n3 bool = false
	var n4 byte = 'a'

	//fmt.Sprintf(“%参数”,表达式)，这个方式比demo12的strconv的方式好用，重点联系这个

	s1 := fmt.Sprintf("%d", n1)
	fmt.Printf("s1对应的类型是: %T, s1 = %q \n", s1, s1) //%q 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

	s2 := fmt.Sprintf("%f", n2)
	fmt.Printf("s2对应的类型是: %T, s2 = %q \n", s2, s2) //%v 值的默认格式表示,所以，都换成%q，其实就是比%v显示的时候多了“”

	s3 := fmt.Sprintf("%t", n3)
	fmt.Printf("s3对应的类型是: %T, s3 = %q\n", s3, s3)

	s4 := fmt.Sprintf("%c", n4) //注意，找相关的api介绍，不能用%s，得用%c
	fmt.Printf("s4对应的类型是: %T, s4 = %q\n", s4, s4)

}
