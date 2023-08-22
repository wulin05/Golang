package main

import (
	"fmt"
)

/*
第六章: 错误处理
第一节：defer + recover机制处理错误
*/

func main() {
	a := test01()
	fmt.Printf("test01函数的除法操作执行成功,返回值为: %v \n", a)
	fmt.Println("正常执行下面的逻辑.....")

	fmt.Println("----------------------------------------------------")
	fmt.Println("test02函数代码有误,调用test02()函数后,通过defer+recover让程序捕获错误,并能继续运行后续代码：")
	test02()
	fmt.Println("test02函数虽然出错,但是被捕获了,所以后续代码还能照样被执行！")

}

func test01() int {
	num1 := 10
	num2 := 2 // 将除数改成0后,程序出现错误/恐慌后,程序被中断,无法继续执行
	result01 := num1 / num2
	return result01
}

func test02() {
	//利用defer + recover来捕获错误: defer后加上匿名函数的调用
	defer func() {
		//调用recover内置函数,可以捕获错误:
		err := recover()
		//如果没有捕获错误,返回值为零值: nil
		if err != nil {
			fmt.Println("错误已经被捕获")
			fmt.Println("err是: ", err)
		}
	}() //注意这个地方必须加上(),相当于匿名函数定义后,直接使用,因为匿名函数没有传入参数,所以就(),如果有传入参数,就相应地添加！
	num3 := 10
	num4 := 0
	result02 := num3 / num4
	fmt.Println(result02)
	/*
		与test02没有使用return语句,所以在main函数中就只要test02()就可以,而test01有返回值,所以
		使用了一个变量a := test01()来接收这个返回值;
		如果在main函数中这样写：b := test02(),程序会报错！
		或者在test02(),写了一个 return result02,那么也是报错！
	*/
}
