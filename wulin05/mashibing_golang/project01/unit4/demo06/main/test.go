package main

import "fmt"

func main() {
	//功能:求1-100的和,当和第一次超过300的时候,停止程序
	var sum int = 0
	for i := 1; i<=100; i++{
		sum += i
		fmt.Println(sum)
		if sum >= 300 {
			//break是停止正在执行的这个循序
			break
		}
	}
	fmt.Println("...........................1")

	//双重循环:
	for m := 1; m <= 5; m++{
		for n := 2; n <= 4; n++{
			fmt.Printf("m: %v, n: %v \n", m, n)
		}
	}

	fmt.Println("---------------------------2")

	//双重循序中break的作用范围:
	for m := 1; m <= 5; m++{
		for n := 2; n <= 4; n++{
			fmt.Printf("m: %v, n: %v \n", m, n)
			if m == 2 && n == 2 {
				break //从运行结果可以看出,break作用是结束离它最近的循环,然后到外层循环继续
			}
		}
	}

	fmt.Println("---------------------------3")
	//使用lable标签来决定作用的范围
	lable1:
	for m := 1; m <= 5; m++{
		//lable2: 这个标签没有使用到,所以要注释掉,否则报 定义未使用 的错
		for n := 2; n <= 4; n++{
			fmt.Printf("m: %v, n: %v \n", m, n)
			if m == 2 && n == 2 {
				break lable1 //结束指定标签对应的循环
			}
		}
	}

	fmt.Println("---------------------------4")

	/*
	总结break
	*/
}
