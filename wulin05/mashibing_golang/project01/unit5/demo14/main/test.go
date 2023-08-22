package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
系统函数: 字符串相关函数
1.统计字符串长度,按字节进行统计: len(str),使用内置函数不用导包,直接用就行
2.字符串遍历: r := []rune(str)
3.字符串转整数：strconv.Atoi
4.整数转字符串：strconv.Itoa
关于字符串和其他基本数据类型之间的转换,可以查看unit2的从demo11、demo12、demo13
*/

func main() {
	//统计字符串的长度,按字节进行统计
	str01 := "golang你好" //在golang中,汉字是utf-8字符集,一个汉字3个字节
	fmt.Println(len(str01))

	//对字符串进行遍历:
	//1. 普通的for循环是按字节进行遍历,遇到中文的话就会出现乱码
	//2. 使用for range的方式,就没问题了
	for i, value := range str01 {
		fmt.Printf("索引为: %d,具体的值为: %c \n", i, value) //%c也可以用%q
	}
	//3. 利用r := []rune(str)
	str02 := "linwu林武"
	fmt.Printf("通过变量名str02直接输出字符串: %q \n", str02) //目前没搞懂,这样不就直接能打印出"linwu林武"了,用下面的切片干卵？
	r := []rune(str02)
	for j := 0; j < len(r); j++ {
		fmt.Printf("%c \n", r[j])
	}

	//字符串转整数
	num1, _ := strconv.Atoi("666")                  //这里注意下,如果字符串不是例子那样的数字,是中文,就会有错误
	num2, _ := strconv.Atoi("A")                    //所以,strconv.Atoi这个函数是有两个返回值的,用"_"来接收错误信息
	fmt.Printf("输出num1=%v, num2=%v \n", num1, num2) //num=0 !!!

	//整数转字符串
	str03 := strconv.Itoa(88)
	fmt.Println(str03)

	//统计一个字符串有几个指定的子串
	count := strings.Count("golangandjavagago", "go")
	fmt.Println(count) // 2

	//不区分大小写的字符串比较
	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)

	//区分大小写进行字符串比较
	fmt.Println("hello" == "Hello")

	//返回子串在字符串第一次出现的索引值,如果没有返回-1:
	fmt.Println(strings.Index("golangandjavagago", "ga")) //第一次出现ga的索引是5

}
