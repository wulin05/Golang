package main

import (
	"fmt"
)

func main(){
    //+加号：
	//1.正数 2.相加操作 3.字符串拼接
	var n1 int = +10
	fmt.Println(n1)
	var n2 int = 4 + 7
	fmt.Println(n2)
	var s1 string = "abc" + "def"
	fmt.Println(s1)

	// /除号
	fmt.Println(10/3) //两个int类型数据运算，结果为int
	fmt.Println(10.0/3) //浮点数类型参与运算，结果为浮点类型

	// %取模 等价于: a%b=a-a/b*b
	fmt.Println(10%3)    //10%3 = 10-10/3*3                等于1
	fmt.Println(-10%3)   //-10%3 = -10-(-10)/3*3           等于-1
	fmt.Println(10%-3)   //10%(-3) = 10-10/(-3)*(-3)       等于1
	fmt.Println(-10%-3)  //-10%(-3) = -10-(-10)/(-3)*(-3)  等于-1

	//++ 自增操作, -- 自减操作
	// go语言里，++, -- 操作非常简单，只能单独使用，不能参与到运算中去
	// go语言里，++, --只能在变量的后面，不能在变量的前面 ++a,--a 是错误写法
	var a int = 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
}