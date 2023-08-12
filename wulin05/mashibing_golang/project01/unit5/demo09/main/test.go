package main

import (
	"fmt"
)

/*
	该代码是为了说明以下内容: 
	1.在go中,函数也是一种数据类型,可以赋值给一个变量,则该变量就是一个函数类型的变量了,
	  通过该变量可以对函数调用

	2.
*/

//定义一个函数test01
func test01 (num int) {
	fmt.Println(num)
}

//定义一个函数test02, 第三个形参名字是 testFunc, 类型是 func(int),如下:
func test02 (num1 int, num2 float32, testFunc func(int)) {
	fmt.Println("-------test02函数")
}

func main() {
	//函数也是一种数据类型,可以赋值给一个变量
	a := test01  //变量a的类型就是test01函数类型
	fmt.Printf("a的类型是: %T, test01函数的类型是: %T \n", a, test01)
	//通过a变量可以对函数test01调用
	a(10) // 等价与 test01(10)

	fmt.Println("-----------------------------------------------------------------1")


	//调用test02函数,将test01函数作为形参
	test02(10, 3.14, test01)
	test02(10, 3.14, a)
	fmt.Printf("test02函数的类型是: %T \n", test02)

	fmt.Println("-----------------------------------------------------------------2")


	//自定义数据类型: (相当于起别名)
	type myInt int           //给int类型起别名叫myInt类型
	var score01 myInt = 99     
	fmt.Println("score01=", score01)

	var score02 int = 88
	fmt.Println("score02=", score02)
	score02 = int(score01) //因为直接 score02 = score01 会报错: 类型不匹配,也就是说myInt和int类型不一样
	fmt.Printf("新的score02的值是: %v \n", score02)

	fmt.Println("-----------------------------------------------------------------3")


	//调用函数test03
	test03(5, "linwu", a)
	test03(5, "linwu", test01)
	//test03(5, "linwu", test02) 这样是不行的,因为test02的类型是 func(int, float32, func(int))
	//                           test03的类型是：func(int, string, func(int))
	//                           那么,就没法将test02作为test03的第三个参数了,类型都不匹配。
	// 除非test03的类型是：func(int, string, func(int, float32, func(int))),那就可以！！
	// 这个很容易混淆哦：即,因为test03的第三个形参的类型是func(int), 而test01函数类型是这个,所以没问题；
	// 但是 test02的类型是 func(int, float32, func(int)), 就不是func(int)类型,所以不行！
	// 除非,这样定义test03函数:
	// type newFunc func(int, float32, func(int))
	// func test03 (num1 int, str string, newFunc)
	// 那就可以这样用了： test03(5, "linwu", test02)
}

	//
	//定义一个函数test03
	type myFunc func(int) // 函数类型func(int),用myFunc来代替
	func test03 (num1 int, str string, testFunc myFunc) {
		fmt.Println("-------test03函数")
	}