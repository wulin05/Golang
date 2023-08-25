package main

import (
	"fmt"
)

/*
方法的注意事项：
3. GO中的方法作用在指定的数据类型上的,和指定的数据类型绑定,因此自定义类型,都可以有方法,而不仅仅是struct,比如int,float32等都可以有方法
*/

type myint int //自定义myint类型与指定的int类型绑定

func (i myint) print() {
	// i = 30  同样i是myint类型,也是值传递,不会对main函数中的原元素造成影响。
	fmt.Println("输出值: ", i)
}

func (j *myint) change() { // 使用这种指针方式,就相当于change方法的j变量是指针,存的地址是main函数中num变量的地址,所以,经过下面 *j = 30,就相当于直接修改num的值
	*j = 30
	fmt.Println("j = ", *j)
}

//上面的两个方法：print,change,是自定义类型myint的方法

func main() {

	var num myint = 20
	num.print() //类型是myint的num变量,调用了print()方法
	// fmt.Println(num)

	num.change() //类型是myint的num变量,调用了change()方法
	fmt.Println(num)
}
