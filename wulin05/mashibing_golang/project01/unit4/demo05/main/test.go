package main

import "fmt"

func main() {
	//定义一个字符串
	var str string = "hello golang你好"
	//方式一: 普通for循环: 按照字节进行便利输出的(暂时先不使用中文)
	// for i := 0; i < len(str); i++{
	// 	fmt.Printf("%c \n", str[i])
	// }

	//方式二: for range
	for i, value := range str {
		fmt.Printf("索引为: %d,具体的值为: %c \n", i, value)
		//对str进行遍历,遍历的每个结果的索引值被i接收,每个结果的具体值被values接收
		//遍历对字符进行遍历的
	}

}
