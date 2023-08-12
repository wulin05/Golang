package main

import "fmt"

func main() {
	//联系转义字符
	fmt.Println("aaa\nbbb")   // \n 换行
	fmt.Println("aaa\bbbb")   // \b 表示退格，在输出bbb，结果就是aaa变成aa，然后在bbb，最终是aabbb
	fmt.Println("aaaaa\rbbb") // \r 表示光标回到开头，然后输出bbb，原先aaaaa，变成了bbbaa

	fmt.Println("aaaaaaaaaaaaaa")
	fmt.Println("aaaaa\tbbbbbb") // \t 表示制表符，每8个字符是一个制表符
	fmt.Println("aaa\tbbb")
	fmt.Println("aaaaaaaa\tbbbbb")

	// \就是转义字符，将后面的字符原样输出
	fmt.Println("\"Golang\"")
}
