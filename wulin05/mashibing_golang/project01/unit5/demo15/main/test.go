package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串的替换: strings.Replace("golangjavagogo","go","LIN",n)
	//关于最后的n参数,可以指定你希望替换几个,如果n=-1表示全部替换,替换两个，n就是2
	str1 := strings.Replace("golangjavagogo", "go", "LIN", -1)
	fmt.Println(str1)
	str2 := strings.Replace("golangjavagogo", "go", "LIN", 2)
	fmt.Println(str2)

	//按照指定的某个字符,为分割标识,将一个字符串拆分成字符串数组
	//strings.Split("go-python-java","-")
	arr := strings.Split("go-python-java", "-")
	fmt.Println(arr)

	//将字符串的字母进行大小写的转换
	str3 := strings.ToLower("GOLANG") //golang
	fmt.Println(str3)
	str4 := strings.ToUpper("golang") //GOLANG
	fmt.Println(str4)

	//将字符串左右两边的空格去掉:
	fmt.Println(strings.TrimSpace("    go and python   ")) //go and python

	//将字符串最左右两边指定的字符去掉：
	fmt.Println(strings.Trim("~golang~ ", "~")) //golang~
	fmt.Println(strings.Trim("~golang~", "-"))  //golang

	//将字符串左边指定的字符去掉：
	fmt.Println(strings.TrimLeft("~golang~", "~")) //golang~

	//将字符串右边指定的字符去掉：
	fmt.Println(strings.TrimRight("~golang~", "~")) // ~golang

	//判断字符串是否以指定的字符串开头:
	fmt.Println(strings.HasPrefix("https://redmica.jcbc.co.jp", "https")) //true

	//判断字符串是否以指定的字符串结束:
	fmt.Println(strings.HasSuffix("https://redmica.jcbc.co.jp", "com")) //false

}
