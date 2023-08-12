package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 18
	var s1 string = strconv.FormatInt(int64(n1), 10) //参数:第一个参数必须转为int64类型，第二个参数指定字面值的进制形式为十进制
	fmt.Printf("s1对应的类型是: %T, s1 = %q \n", s1, s1)

	var n2 float32 = 3.14159267
	//第二个参数:'f'代表(ddd.dddd) 第三个参数: 8 保留小数点后面8位 第四个参数: 表示这个小数是float32类型
	var s2 string = strconv.FormatFloat(float64(n2), 'f', 9, 32)
	//这个strconv.FormatFloat函数的第一个参数要求必须是float64类型，所以.....
	fmt.Printf("s2对应的类型是: %T, s2 = %q \n", s2, s2)

	var n3 float64 = 3.14159267
	var s3 string = strconv.FormatFloat(n3, 'f', 9, 64)
	fmt.Printf("s3对应的类型是: %T, s3 = %q \n", s3, s3)

	var n4 bool = true
	var s4 string = strconv.FormatBool(n4)
	fmt.Printf("s4对应的类型是: %T, s4 = %q \n", s4, s4)
}
