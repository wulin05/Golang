package main

import (
	"fmt"
)

/*
方法的注意事项：
4. 方法的访问范围控制的规则,和函数一样。方法名首字母小写,只能在本包访问,方法首字母大写,可以在本包和其它包访问
5. 如果一个类型实现了String()这个方法,那么fmt.Println默认会调用这个变量的String()进行输出
   以后定义结构体的话,最好就定义String()作为输出结构体信息的方法, fmt.Println会自动调用
*/

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string { // 使用结构体Student的指针类型作为参数类型, 方法名:String, 输出结果是一个字符串:string
	str := fmt.Sprintf("Name = %v, Age = %v", s.Name, s.Age)
	return str
}

func main() {
	stu := Student{
		Name: "乐乐",
		Age:  18,
	}
	fmt.Println(stu) // {乐乐 18}

	//传入地址,如果绑定了String方法就会自动调用
	// fmt.Println(stu) 这样也行,不报错。
	fmt.Println(&stu) // Name = 乐乐, Age = 18

	//可以这样调用么？当然可以,因为有返回值,所以用A来接收返回值,在打印A。
	A := stu.String()
	fmt.Println(A) // Name = 乐乐, Age = 18

}
