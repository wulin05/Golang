package main

import "fmt"

/*
内置函数：
1. golang设计者为了编程方便,提供了一些函数,这些函数不用导包就可以直接使用,我们称为Go的内置函数/内建函数
2. 在builtin包下,使用内置函数也行,直接用就行
3. 常用函数:
(1)len函数:
统计字符串的长度,按字节进行统计
(2)new函数：
分配内存,主要用来分配值类型(int系列,float系列,bool,string,数组和结构体struct)
(3)make函数: 关于make函数,由于slice切片、map等还没学,后续介绍
分配内存,主要用来分配引用类型(指针、slice切片、map、管道chan、interface等)
*/

func main() {
	//len函数
	str1 := "golang"
	fmt.Println(len(str1)) //len函数不需要导入buildin包,直接调用就可以了
	str2 := "python林武"
	fmt.Println(len(str2))

	//new函数分配内存,new函数的实参是一个类型而不是具体数值,new函数返回值是对应类型的指针: num: *int
	num := new(int)
	fmt.Printf("num的类型是: %T, num的值是: %v \n", num, num) //num的类型是: *int, num的值是: 0xc00001a0e0
	fmt.Printf("num自己本身的地址是: %v, num指针指向的地址所对应的值是: %v", &num, *num)
}
