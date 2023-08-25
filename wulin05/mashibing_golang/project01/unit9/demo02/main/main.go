package main

import (
	"fmt"
)

/*
映射(集合)：
1. map的操作：增加、修改、删除、清空、查找、获取长度、遍历
*/

func main() {
	//定义一个map
	var a map[int]string = make(map[int]string)
	//或者 a := make(map[int]string)
	//操作一：增加
	a[20230801] = "张三"
	a[20230802] = "李四"
	a[20230803] = "王五"
	fmt.Println(a)

	//操作二：修改
	a[20230801] = "刘六"
	fmt.Println(a)

	//操作三：删除、delete
	delete(a, 20230801)
	fmt.Println(a)

	fmt.Println("--------------------------------")
	//操作四：清空map
	//方式一： 使用遍历方式，遍历key，逐个删除
	for i, _ := range a {
		delete(a, i)
	}
	fmt.Printf("a集合使用for循环遍历清空元素: %v\n", a)

	fmt.Println("--------------------------------")
	//方式二： make一个新的，让原来的成为垃圾，被gc回收
	//重新定义一个map(集合)b
	var b map[int]string
	b = make(map[int]string, 8)
	b[20230901] = "王朝"
	b[20230902] = "马汉"
	b[20230903] = "张龙"
	fmt.Println("b集合元素:", b)

	b = make(map[int]string)
	fmt.Printf("使用make清空b集合: %v\n", b)

	fmt.Println("--------------------------------")
	//操作五：查找
	// value,bool = map[key]: value为分支的value,bool为是否返回,要么true要么false
	//在定义一个map：c
	var c map[int]string = map[int]string{
		20231001: "小美",
		20231002: "小明",
	}
	value, flag := c[20231002]
	fmt.Println(value)
	fmt.Println(flag)

}
