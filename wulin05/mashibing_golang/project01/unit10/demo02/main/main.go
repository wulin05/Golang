package main

import (
	"fmt"
)

/*
结构体之间的转换：
1. 结构体是用户单独定义的类型,和其他类型进行转换时需要有完全相同的字段(名字、个数和类型)
2. 结构体进行type重新定义(相当于取别名),GO认为是新的数据类型,但是相互间可以强转
*/

type Student struct {
	Age int
}

type Stu Student //重新定义Student类型为Stu。结构体进行type重新定义(相当于取别名)，GO任务是新的数据类型,但是相互间可以强转！如 2 的例子：直接赋值不行,要先转了才行。

type Person struct {
	Age int
}

func main() {

	//1. 结构体是用户单独定义的类型,和其他类型进行转换时需要有完全相同的字段(名字、个数和类型)
	var s Student = Student{} //这就是demo01的方式 2.2,只是这边value只有一个
	var p Person = Person{30}
	s = Student(p) //注意不能直接 s = p 方式将p的值赋给s,要先将p的类型由 Person结构体类型强制转为Student类型
	fmt.Println(s)

	//2. 结构体进行type重新定义(相当于取别名),GO认为是新的数据类型,但是相互间可以强转
	var s1 Student = Student{}
	var s2 Stu = Stu{99}
	// s1 = s2  这样会报错,类型不匹配！
	s1 = Student(s2)
	fmt.Println(s1)

}
