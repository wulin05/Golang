package main

import (
	"fmt"
)

/*
方法的注意事项：
1. 结构体类型是值类型,在方法调用中,遵守值类型传递机制,是值拷贝传递方式，也叫副本拷贝(就是前一章节说的值传递,不会影响结构体的原数据)。
2. 如希望在方法中,改变结构体变量的值,可以通过结构体指针的方式来处理
*/

type Student struct {
	Id   int
	Name string
}

func (s *Student) change() {
	s.Id = 380          //其实真正应该写成 (*s).Id = 380
	s.Name = "武林"       //同理真正应该写成 (*s).Name = "武林"
	fmt.Println(s.Id)   //同理应该写成 fmt.Println((*s).Id)
	fmt.Println(s.Name) //同理应该写成 fmt.Println((*s).Name)
} //原因：底层编译器做了优化,自动帮加上 &、*

func main() {
	//初始一个结构体n1,并赋值
	var m Student = Student{38, "林武"}

	fmt.Printf("m的内存地址是: %p", &m)
	m.change() //其实真正应该写成 (&m).change(),因为change方法的传入参数是指针: *Student

	fmt.Println(m.Id)   //跟上面一样,应该是fmt.Println((*m).Id)
	fmt.Println(m.Name) //同理也原因, fmt.Println((*m).Name)

	//总的一句话,所以* &省了,程序也不报错！
}
