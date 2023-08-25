package main

import (
	"fmt"
)

/*
创建结构体实例时指定字段值
其实就是demo01的延伸
1.方式一：按照顺序赋值
2.方式二：按照指定类型
3.方式三：想要返回结构体的指针类型
*/

//定义结构体
type Student struct {
	Name string
	Age  int
}

func main() {
	//方式一：按照顺序赋值操作
	var s1 Student = Student{"小李", 22}
	fmt.Println(s1)

	//方式二：按照指定类型
	s2 := Student{
		Name: "小王",
		Age:  21,
	}
	fmt.Println(s2)

	//方式三：返回结构体的指针类型
	var s3 *Student = &Student{"小林", 20}
	fmt.Println(s3)  // &{小林 20}
	fmt.Println(*s3) // {小林 20}
	fmt.Println(&s3) // 0xc0000ca020
}
