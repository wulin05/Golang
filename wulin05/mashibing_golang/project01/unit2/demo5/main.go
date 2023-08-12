package main

import "fmt"

func main() {
	//定义字符类型的数据:
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '6'
	fmt.Println(c2)
	var c3 byte = '('
	fmt.Println(c3)
	//字符类型，本质上就是一个整数，也可以直接参与运算，输出字符的时候，会将对应的码值做一个输出
	fmt.Println(c3 + 20)
	//字母、数字、标点等字符，底层是按照ASCII进行存储

	//var c4 byte = '中'   这个是错误的，汉字字符，底层对应的是Unicode码值
	//对应的码值为20013,byte类型溢出，能存储的范围：可以用int
	var c4 int = '中'
	fmt.Println(c4)
	//总结：Golang的字符对应的使用的是UTF-8编码
	//这么说：ASCII和Unicode都是字符集，ASCII是Unicode字符集的前128位，而utf-8是字符集的编码方案
	//即，utf-8是Unicode的其中一个编码方案，还有utf-16,utf-32,utf-64 编码 。

	var c5 byte = 'A'
	var c6 int = '国'
	//想显示对应的字符，必须采用格式化输出
	fmt.Printf("c5对应的具体的字符为: %c", c5)
	fmt.Println()
	fmt.Printf("c6对应的具体的字符为: %c", c6)

}
