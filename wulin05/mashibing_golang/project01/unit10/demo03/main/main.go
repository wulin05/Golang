package main

import (
	"fmt"
)

/*
方法的引入：
1. 方法是作用在指定的数据类型上,和指定的数据类型绑定,因此自定义类型,都可以有方法,而不仅仅是结构体struct
2. 方法的声明和调用格式
type A struct {         //这是结构体A
	Num int
}
func (a A) test() {     //这是方法的声明,含义是：结构体A的方法test()。注意：如果把(a A)去掉,变成 func test() {} ,这样就是函数了,所以记住方法和函数的区别
	fmt.Println(a.Num)
}
调用：
var a A
a.test()

(1) func (a A) test()相当于A结构体有一个方法,叫test
(2) (a A)体现方法test和结构体A绑定关系：方法是作用在指定的数据类型上,和指定的数据类型绑定
*/

//定义Person结构体
type Person struct {
	Name string
}

//给Person结构体绑定方法test：其实可以这么理解,就相当于结构体有了test()方法。
//参数名字随便起,p可以是任何其他的,只是习惯是结构体的首字母小写,可以写成 func (m Person) test(){}
func (p Person) test() {
	fmt.Println(p.Name)
}

func main() {
	//创建结构体对象:
	var p Person //同理,这里的p不是说要跟上面的方法 func (p Person) test()里的p要一致,只是一个变量名,可以写成 var n Person
	p.Name = "林武"
	p.test()

	var m Person
	m.Name = "武林"
	m.test()
	//从上面两个结构体调用test方法可以看出,结构体对象传入test方法中,是值传递,和函数参数传递一致。

	//如果其他类型变量调用test方法一定会报错,例如：
	//定义一个不是结构体类型的变量 n,比如类型是map,这个n就不能调用test方法,因为test方法已经和Person这个结构体类型绑定了!
	// n := make(map[int]string)
	// n.test()   -----> 报错: .\main.go:48:4: n.test undefined (type map[int]string has no field or method test)

	//继续延伸:
	//如果在test()方法中,将p或m结构体中Name值修改了,那么在func main()函数中p或m的Name值会改变吗？
	//由于上面的改了怕到时候再看有点乱,举例的话,我就重新定义一个同样是绑定Person结构体的方法：change()
	//通过最后的两个输出,得出结论：因为结构体也是值传递(也叫副本传递),不是引用传递(想象下指针),通过内存分析,change()方法是对传入的Name值修改,不会改变结构体原Name的值
	var j Person
	j.Name = "偌小王"
	j.change()          //王小偌
	fmt.Println(j.Name) //偌小王
}

func (q Person) change() {
	q.Name = "王小偌"
	fmt.Println(q.Name)
}

//可以认为给结构体Person又添加了一个方法change(); 结构体Person有两个方法: test() 和 change()
