package main

import (
	"fmt"
)

/*
映射(集合)：
接上一章节：map的操作
5.获取mpa长度
6.遍历: for range
*/

func main() {
	//定义一个map
	a := map[int]string{
		20230101: "语文",
		20230102: "数学",
		20230103: "物理",
	}
	a[20230104] = "化学"
	fmt.Println(a)

	//获取长度:
	fmt.Println(len(a))

	fmt.Println("-------------------------------------------------------------------------------")
	//遍历:
	for k, v := range a {
		fmt.Printf("a[%v] = %v \n", k, v)
	}

	fmt.Println("-------------------------------------------------------------------------------")

	//加深难度: map中的value也是map类型
	b := make(map[string]map[int]string, 3)
	//赋值操作
	b["班级1"] = map[int]string{
		20060701: "林武",
		20060702: "黄鑫杰",
		20060703: "陈何峰",
	}
	// b["班级1"] = ([20060704] = "冯晓强")  这样添加元素是错误的,应该如下
	//下面不是添加元素,变成 20060704: "冯晓强" 把原键值对给覆盖了：map[班级1:map[20060704:冯晓强]]
	//下面的方式相当于上一章节的清空操作,如果{}里面没有内容的话,不同的是：{}是将旧集合的value赋为空！
	// b["班级1"] = map[int]string{
	// 	20060704: "冯晓强",
	// }

	//应该这样：
	b["班级1"][20060704] = "冯晓强" //另外需要说明的是:虽然定义b的时候,定义了长度为3,但不必担心超过容量的问题
	//go语言的map在内部会自动处理存储的扩展和管理
	fmt.Println(b)

	b["班级2"] = make(map[int]string)
	b["班级2"][20060801] = "丽丽"
	b["班级2"][20060802] = "萌萌"
	b["班级2"][20060803] = "乐乐"
	fmt.Println(b)

	//遍历:
	for className, students := range b {
		fmt.Printf("%s{", className)
		for studentNo, name := range students {
			fmt.Printf("%d:%s\t", studentNo, name)
		}
		fmt.Printf("}\n")
	}
}
