package main

import (
	"fmt"
)

/*
方法和函数的区别

3. 对于函数来说,参数类型对应是什么就要传入什么
4. 对于方法来说,接收者为值类型,可以传入指针类型,接收者为指针类型,可以传入值类型：因为底层编译器做了优化
*/

//定义结构体
type Student struct {
	Name string
	Age  int
}

//定义函数methd01
func method01(i Student) {
	fmt.Println(i.Name)
}

//定义函数method02
func method02(j *Student) {
	fmt.Println((*j).Age)
}

//定义方法test01,test02
func (m Student) test01() {
	fmt.Println(m.Name)
}

func (n *Student) test02() {
	fmt.Println((*n).Age)
	// fmt.Println(n.Age) 这样也可以,不会报错
}

func main() {

	//创建Student结构体实例stu
	stu := Student{
		Name: "王佳华",
		Age:  30,
	}

	//调用函数method01和method02,用于说明：3. 对于函数来说,参数类型对应是什么就要传入什么
	method01(stu)
	//method01(&stu) 报错
	method02(&stu)
	//method02(stu) 报错

	//调用方法,用于说明：4. 对于方法来说,接收者为值类型,可以传入指针类型,接收者为指针类型,可以传入值类型：因为底层编译器做了优化
	//unit10的demo04也有这方面的介绍
	(&stu).test01() //虽然使用&stu,但是也不会报错,因为底层会自动使用值传递
	stu.test01()
	//方法test02()正常应该用第一种,但是第二种也不会报错
	(&stu).test02()
	stu.test02()

}
