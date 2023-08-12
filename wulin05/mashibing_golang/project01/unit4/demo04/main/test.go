package main

import "fmt"

func main() {
	/*
	for循环的语法格式:
	for 初始表达式; 布尔表达式; 迭代因子{
		循环体;
	}

	注意: for的初始表达式 不能用var定义变量的形式, 要用 := 的方式
	注意: for循环实际就是让程序员写代码效率提高, 代码简洁了, 但是底层该怎么执行还是怎么执行的, 底层效率并没有提高
	*/

	//方式一:
	var sum int = 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	//方式二: 
	j := 1
	for j <= 5 {
		fmt.Println("我爱golang !")
		j++
	}

	/*
	死循环一:
	for {
	 	fmt.Println("你好,Golang !")
	}

	死循环二:
	for ;; {
		fmt.Println("你好,Golang !")
	}
	*/

}
