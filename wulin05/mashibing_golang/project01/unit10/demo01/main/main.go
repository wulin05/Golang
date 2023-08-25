package main

import (
	"fmt"
)

/*
面向对象：GO支持面向对象编程(OOP)
1.面向对象的引入: GO语言和传统的面向对象变成有区别,并不是纯粹的面向对象语言,所以应该说GO支持面向对象变成特性比较准确
2.GO没有类(class),GO语言的结构体(struct)和其它编程语言的类(class)有同等的地位,可以理解GO是基于struct来实现OOP特性
3.GO面向对象编程非常简洁,去掉传统OOP语言的方法重载、构造函数和析构函数、隐藏的this指针等等
4.GO仍然有面向对象编程的继承、封装和多态的特性。只是实现的方式和其它OOP语言不一样,比如继承：GO语言灭有extends关键字,
  继承是通过匿名字段来实现
*/

//定义老师结构体,将老师中的各个属性 统一放入结构体中管理
type Teacher struct {
	Name   string //变量名字大写外界才可以访问这个属性
	Age    int
	School string
}

func main() {
	//创建老师结构体的实例、对象、变量, 即,结构体实例的创建方式
	//方式 1：
	var t1 Teacher
	fmt.Println(t1) //在未赋值时默认值: { 0 }, 注意 { 0 } 内0的两边是空字符""
	t1.Name = "林武"
	t1.Age = 37
	t1.School = "FZU"
	fmt.Println(t1)              //{林武 37 fzu}
	fmt.Printf("t1的类型是: %T", t1) //t1的类型是: main.Teacher

	//方式 2.1：在视频中,方式二和方式三其实是一种！！
	var t2 Teacher = Teacher{}
	fmt.Println(t2) // 啥东西都没显示！！
	t2.Name = "王小诺"
	t2.Age = 35
	t2.School = "HY"
	fmt.Println(t2)

	//方式 2.2:
	var t3 Teacher = Teacher{"王佳华", 39, "FQ"}
	fmt.Println(t3)

	//方式 3：
	var t4 *Teacher = new(Teacher) //t4是指针,其实指向的就是地址,应该给这个地址指向的对象的字段赋值
	(*t4).Name = "陈丹婷"             //*的作用: 根据地址取值
	(*t4).Age = 37
	t4.School = "CT" //为了符合程序员编程习惯,go语言提供简化的赋值方式,但其实编译器底层对t4.School
	//还是进行了转换 (*t4).School = "CT"
	fmt.Println(*t4)

	//方式 4:
	var t5 *Teacher = &Teacher{"张燕", 37, "HF"}
	fmt.Println(*t5)
}
