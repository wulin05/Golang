package main

import (
	"fmt"
)

/*
映射：
1. map的引入：map是Go语言中内置的一种类型,它将键值相关联,可以通过键key来获取对应的值value。类似于其他语言的集合
2. 基本语法: map[keytype]valuetype
   key、value的类型: bool、数字、string、指针、channel、还可以是只包含前面几个类型的接口、结构体、数组
   key通常为int、string类型；value通常为数字(整数、浮点数)、string、map、结构体
   key部分：slice、map、function不可以
3. map的创建方式：三种
*/

func main() {
	//方式1：
	//声明一个map的变量
	var a map[int]string //波浪线的原因是vscode建议将变量声明和赋值合并到一行
	//只声明map内存是没有分配空间
	//必须通过make函数进行初始化,才会分配空间
	a = make(map[int]string, 10) //map可以存放10个键值对
	//将键值对存入map中：
	a[20230801] = "张三"
	a[20230802] = "李四"
	a[20230803] = "王五"
	// a[20230803] = "朱六" //key不可以重复的,如果遇到重复,后一个key的value会替换前一个同一个key的value
	a[20230804] = "张三" // value可以重复

	//输出映射、也可以叫集合
	fmt.Println(a) // map[20230801:张三 20230802:李四 20230803:王五 20230804:张三]
	// 从输出来看,map存放的数据是无序的

	//方式2：
	b := make(map[int]string)
	b[20230901] = "小明"
	b[20230902] = "小美"
	fmt.Println(b)

	//方式3：
	c := map[int]string{
		20231001: "王朝",
		20231002: "马汉",
	}
	c[20231003] = "赵虎"
	fmt.Println(c)

}
