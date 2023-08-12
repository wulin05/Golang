package main

import "fmt"

//基本数据类型和数组默认都是值传递的,即进行值拷贝。在函数内修改,不会影响到原来的值
//以值传递方式的数据类型,如果希望在函数内的变量能修改函数外的变量,可以传入变量的地址&,
//函数内以指针的方式操作变量,从效果来看类似引用传递


//从下面的main函数调用test函数的结果,可以看出,基本数据类型和数组的这种方式是值传递,
//在test函数中是改变不了main函数中的m的值: 请脑海里显现出栈帧、堆、代码区的图形！！！！！
func test (num int) {
	num = 30
	fmt.Println("test-----", num)
}

func main() {
	var m int = 10
	test(m)
	fmt.Println("main-----", m)
}

