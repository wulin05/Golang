package main

import "fmt"

func main() {
	/*
	基本语法:
    if 条件表达式1 {
		逻辑代码1
	} else if 条件表达式2 {
		逻辑代码2
	}
	......
	else {
		逻辑代码n
	}
	
	实现功能:根据给出的学生的分数,判断学生的成绩:
	>=90 ----->优秀
	>=80 ----->良好
	>=70 ----->一般
	>=60 ----->及格
	<60  ----->垃圾
	*/

	//方式一: 利用if单分支实现:
	fmt.Println("请输入学生的成绩:")
	var score01 int
	fmt.Scanln(&score01)
	if score01 >= 90 {
		fmt.Println("学生的成绩是优秀!")
	}
	if score01 >= 80 && score01 <90 {
		fmt.Println("学生的成绩是良好!")
	}
	if score01 >= 70 && score01 < 80{
		fmt.Println("学生的成绩是一般!")
	}
	if score01 >= 60 && score01 <70 {
		fmt.Println("学生的成绩是及格!")
	}
	if score01 < 60 {
		fmt.Println("学生的成绩是垃圾!")
	}
	//上面方式一是利用多个单分支拼凑出多个选择,每个分支都是并列的,依次从上到下顺序执行,
	//即使走了第一个分支,那么其他分支也是需要判断的!!! 效率有点低....

	//方式二: 利用if多分支,优点:如果已经走了一个分支了,那么下面的分支就不会再去判断执行了
	if score01 >= 90 {
		fmt.Println("学生的成绩是优秀！")
	} else if score01 >= 80 { //else隐藏: score01 < 90
		fmt.Println("学生的成绩是良好！")
	} else if score01 >= 70 { //score01 < 80
		fmt.Println("学生的成绩是普通！")
	} else if score01 >= 60 { //score01 < 70
		fmt.Println("学生的成绩是及格！")
	} else { //score < 60
		fmt.Println("学生的成绩是垃圾！")
	} //建议else的存在,只有有了else才会真正起到多选一的效果

}
