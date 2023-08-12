package main

import "fmt"

func main() {
	//实现功能: 如果口罩的库存小于30个,提示:库存不足
	fmt.Println("示例一: 请输入库存剩余数量:")
	var count01 int
	fmt.Scanln(&count01) //通过屏幕输入的方式获取库存数量
	// var count01 int = 20
	if count01 < 30 {
		fmt.Println("口罩存量不足")
	}else{ //这个else,其实就是隐含了count01 >= 30
		fmt.Println("口罩库存充足")
	}
    //if后面表达式,返回结果一定是true或false
	//表达式返回结果是ture的话,那么{}中的代码就会执行,反之{}中的代码就不会执行
    //if后面一定要有空格,和条件表达式分隔开来
	//{}一定不能省略
	//条件表达式左右的()是建议省略的
    //在golang里,if后面可以并列加入变量的定义,如下:
	if count02 := 50;count02 < 30 {
		fmt.Println("口罩存量不足!")
	} else{
		fmt.Println("口罩存量充足!")
	}
	//注意: else不能写成下面的方式,下面的方式是错误的：
	// if count := 50;count < 30 {
	// 	fmt.Println("对不起,口罩存量不足!")
	// }
	// else{
	// 	fmt.Println("存量充足!")
	// }

	fmt.Println("示例二: 请输入库存剩余数量:")
	var count03 int
	fmt.Scanf("%d",&count03)
	if count03 >= 30 {
		fmt.Println("口罩库存充足!!")
	} else {
		fmt.Println("口罩库存不足!!")
	}

}
