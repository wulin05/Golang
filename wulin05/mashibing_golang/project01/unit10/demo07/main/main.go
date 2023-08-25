package main

import (
	"fmt"
)

/*
方法和函数的区别
1. 方法需要绑定指定类型,函数不需要
2. 调用的方法不一样:
  (1) 函数的调用方式：
          函数名(实参列表)

  (2) 方法的调用方式:
          变量名.方法名(实参列表)
*/

//定义结构体
type Student struct {
	Name string
	Age  int
}

//定义方法
func (s Student) test01() {
	fmt.Println(s.Name)
}

//定义函数
func method01(s Student) {
	fmt.Println(s.Age)
}

func main() {

	//创建Student结构体实例
	stu := Student{
		Name: "王小偌",
		Age:  30,
	}

	//调用函数
	method01(stu)
	//调用方法
	stu.test01()
}
