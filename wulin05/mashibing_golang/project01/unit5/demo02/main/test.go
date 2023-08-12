package main

import "fmt"

/*
函数细节的介绍：
函数和函数是并列关系,所以我们定义的函数不能写到main函数中去
基本语法:
func 函数名(形参列表)(返回值类型列表) {
	执行语句...
	return + 返回值列表
}

一、函数名：
1.遵循标识符命名规范: 见名知意、驼峰命名
2.首字母不能是数字
3.首字母大写该函数可以被本包文件和其他包文件使用(类似public)
4.首字母小写只能被本包文件使用,其他包文件不能使用(类似private)

二、形参列表: cal (num1 int, num2 int)
形参列表：个数可以是一个,可以是n个参数,可以是0个参数
作用: 接收外来的数据: 
实际参数: 实际传入的数据,比如 cal(10, 20)

三、返回值类型列表: 函数返回值对应的类型应该写在这个列表中
返回0个: 可以删除返回值类型,或者只保留()
返回1个: 如果返回值只有一个的话,括号(int)可以省略括号 int
返回多个: 
*/

//自定义函数: 功能: 两个数相加
func cal1 (num1 int, num2 int) () { //没有返回值,所以这个返回值类型(int)的int要删除掉,或者(int)整个删除也可以
	var sum int = 0
	sum += num1
	sum += num2
	//return sum
	fmt.Println(sum)
}

//自定义函数: 计算两个数的和,两个数的差
func cal2 (num1 int, num2 int) (int, int) { //
	var sum int = 0
	sum += num1
	sum += num2

	var sub int = num1 - num2
	return sum, sub
}


func main() {
	//就不需要像之前那么需要用一个变量去接受cal函数的返回值了
	//因为cal函数里就有打印输出语句了
	//	sum1 := cal(10, 20)
	//  fmt.Println(sum1)
	cal1(10, 20)

	sum1, sub1 := cal2(50, 30)
	fmt.Printf("两数相加的值是: %v, 两数相减的值: %v \n", sum1, sub1)

	sum2, _ := cal2(50, 30)
	fmt.Println(sum2)
}
