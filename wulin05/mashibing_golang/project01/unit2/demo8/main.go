package main

import "fmt"

func main() {
	//1.字符串就是一串固定长度的字符连接起来的字符序列
	//定义一个字符串
	var s1 string = "您好，我喜欢Golang"
	fmt.Println(s1)

	//2.字符串是不可变的：指的是字符串一旦定义好，其中的字符的值不能改变
	var s2 string = "abc"
	fmt.Println(s2)
	s2 = "def"
	fmt.Println(s2)
	//s2[0] = "d"  就是指上面说的其中的字符的值不能改变

	//3.字符串的表现形式：
	//(1).如果字符串中没有特殊字符，字符串的表示形式用双引号括起来就可以
	var s3 string = "我是福清人"
	fmt.Println(s3)
	//(2).如果字符串中有特殊字符，那么字符串的表示形式就要用反引号：``
	var s4 string = `
    package main

	import "fmt"

	func main() {
		//布尔类型也叫bool类型，bool类型数据只允许取值true和false
		//布尔类型占一个字节
		//布尔类型适用于逻辑运算，一般用于程序的流程控制
		//测试布尔类型的数值
		var flag01 bool = true
		fmt.Println(flag01)
	
		var flag02 bool = false
		fmt.Println(flag02)
	
		var flag03 bool = 5 < 9
		fmt.Println(flag03)
		}
    `
	fmt.Println(s4)

	//4.字符串的拼接效果
	var s5 string = "abc" + "def"
	fmt.Println(s5)
	s5 += "ghi"
	fmt.Println(s5)

	//当一个字符串过长的时候：其实就是必须把 + 号保留在上一行末尾，没有 + 号，Golang默认是在尾部隐性加了 \n 换行
	var s6 string = "abc" + "def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" + "def"
	fmt.Println(s6)
}
