package main

import (
	"fmt"
	"github.com/wulin05/mashibing_golang/project01/unit10/demo10/model"
)

/*
跨包创建结构体实例:导包、结构体那边实例名一定首字母大写才可以被其它包访问(引用)
如果结构体实例名首字母是小写,就不行了。
解决：结构体首字母小写,也能被跨包访问 ----->引入 "工厂模式" 概念
*/

func main() {
	//跨包创建结构体Student实例：
	var s model.Student = model.Student{Name: "丽丽", Age: 10} //如果没有导包的话,会报错：.\main.go:13:8: undefined: Student
	fmt.Println(s)

	//上面的方式比较繁琐
	t := model.Student{Name: "美美", Age: 22}
	fmt.Println(t)

	//解决：结构体首字母小写,也能被跨包访问 ----->引入 "工厂模式" 概念
	z := model.NewTeacher("林武", "男", 37)
	fmt.Println(*z)
}
