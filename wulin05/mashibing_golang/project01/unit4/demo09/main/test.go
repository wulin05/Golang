package main

import "fmt"

func main() {
	//关于return的用法,这只是其中之一,后续再继续补充
	for i := 1; i <= 100; i++{
		fmt.Println(i)
		if i == 14 {
			return //结束当前的函数,下面的fmt.Println的语句都不输出了!
		}		   //注意与break区别,break只是结束当前循环,下面的fmt.Println语句还是输出的。
	}
	fmt.Println("hello golang")
}
