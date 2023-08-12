package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string--->bool
	var s1 string = "true"
	var b bool
	//ParseBool这个函数有两个返回值: (value bool, err error
	b, _ = strconv.ParseBool(s1)
	//其中b = %t,t表示值是bool类型,如果是用 %v，表示值的默认格式表示
	fmt.Printf("b的类型是: %T, b = %t \n", b, b)

	//string--->int64
	var s2 string = "37"
	var num1 int64
	num1, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("num1的类型是: %T, num1 = %v \n", num1, num1)

	//string--->float32/float64
	var s3 string = "3.1415"
	var f1 float32
	var f2 float64 = float64(f1)
	f2, _ = strconv.ParseFloat(s3, 64)
	//为什么还使用了f2的变量，因为strconv.ParseFloat函数第一个参数传入值必须是float64
	fmt.Printf("f2的类型是: %T, f2 = %v \n", f2, f2)

	//注意：当字符串是golang,明显不是true或false,转为bool类型,没有意义
	//虽然不报错,但是转成bool类型时,是输出默认值:false
	//同理,字符串是golang,明显不是数值,转为int或float类型,默认输出是: 0
	//总结就是：string向基本数据类型转换的时候，一定要确保string类型能够转成有效的数据类型，否则最后得到的结果就是按照对应的类型的默认值输出
	var s4 string = "golang"
	var c bool
	var d int64
	var e float64
	c, _ = strconv.ParseBool(s4)
	d, _ = strconv.ParseInt(s4, 10, 32)
	e, _ = strconv.ParseFloat(s4, 10)
	fmt.Printf("c的类型是: %T, c = %t \n", c, c)
	fmt.Printf("d的类型是: %T, d = %v \n", d, d)
	fmt.Printf("e的类型是: %T, e = %v \n", e, e)

}
