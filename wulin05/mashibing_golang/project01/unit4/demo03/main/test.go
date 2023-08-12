package main

import "fmt"

func main() {
	/*
	基本语法:
	switch 表达式 {
		case 值1,值2,.....
			语句块1
		case 值3,值4,.....
			语句块2
		....
		default:
			语句块
	}

	同样实现功能: 根据给出的学生分数,判断学生的等级:
	>=90 ----->优秀
	>=80 ----->良好
	>=70 ----->一般
	>=60 ----->及格
	<60  ----->垃圾
	*/

	//标准写法:
	var score01 int
	fmt.Println("请输入score01的值:")
	fmt.Scanf("%d",&score01)
	// //switch后面是一个表达式,这个表达式的结果依次跟case进行比较,满足结果的话就执行冒号后面的代码。
	// //default是用来"兜底"的一个分支,其他case分支都不走的情况下就会走default分支
	// //default分支可以放在任意位置上,不一定非要放在最后。
	switch score01/10 {
		case 10 :
			fmt.Println("学生的等级是完美!")
		case 9 :
			fmt.Println("学生的等级是优秀!")
		case 8 :
			fmt.Println("学生的等级是良好!")
		case 7 :
			fmt.Println("学生的等级是一般!")
		case 6 :
			fmt.Println("学生的等级是及格!")
		case 5, 4, 3, 2, 1, 0 :
			fmt.Println("学生的等级是垃圾!")
		default :
			fmt.Println("学生的成绩有误!")
	}

	fmt.Println()



	//下面这种方式比较少用: switch后不带表达式的方式,当做if分支来使用
	var score02 int32
	fmt.Println("请输入score02的值:")
	fmt.Scanf("%d",&score02)
	switch { 
		case score01 == 1 :
			fmt.Println("aaa")
		case score01 == 2 :       
			fmt.Println("bbb")
		default :
			fmt.Println("why???")
	}

	
	//switch后也可以直接声明/定义一个变量, 但记住要用分号;结束, 不推荐这么写:
	switch score03 := 95; { //这里强调的是: ;号别忘记了
		/* 如果是下面的两种写法也是可以,区别在case后是用值,而不是表达式了,不过同样是不推荐:
		下面的两种写法其实不是定义变量,是说明在switch后可以有多个表达式的方式:
		1. switch fmt.Scanln(&score02); score02/10 {
		2. switch score02 := 95; score02/10 {
			  case 10 :
		        fmt.Println("学生等级为完美!")
		      ...... 
		*/
		case score03 >= 90 :
			fmt.Println("学生等级为优秀!")
			fallthrough   //switch穿透,利用fallthrough关键字,如果在case语句块后增加fallthrough,则会继续执行下一个case,也叫switch穿透
		case score03 >= 80 :
			fmt.Println("学生等级为良好!")
		case score03 >= 70 :
			fmt.Println("学生等级为普通!")
		case score03 >= 60 :
			fmt.Println("学生等级为及格!")
		default :
			fmt.Println("学生等级为垃圾!")
	}


	//switch后多个表达式的写法:
	var score04 int
	fmt.Println("请输入score04的值:")
	switch fmt.Scanln(&score04); score04/10 {
		case 10 :
			fmt.Println("学生的等级是完美!")
		case 9 :
			fmt.Println("学生的等级是优秀!")
		case 8 :
			fmt.Println("学生的等级是良好!")
		case 7 :
			fmt.Println("学生的等级是一般!")
		case 6 :
			fmt.Println("学生的等级是及格!")
		case 5, 4, 3, 2, 1, 0 :
			fmt.Println("学生的等级是垃圾!")
		default :
			fmt.Println("学生的成绩有误!")
	}


	// //这里主要是说明case后的各个值的数据类型,必须和switch后的表达式数据类型一致
	var score05 float32
	fmt.Println("请输入score05的值:")
	fmt.Scanln(&score05)
	switch int(score05 / 10) { //score05是浮点类型,例如95.5/10 = 9.55,case里没有匹配的,最后落入default语句块了!
		case 10:                   //所以在这里就需要强制转换为int类型,这样才能匹配下面的case后面的值的类型
			fmt.Println("学生等级为完美!")
		case 9:
			fmt.Println("学生等级为优秀!")
		case 8:
			fmt.Println("学生等级为良好!")
		case 7:
			fmt.Println("学生等级为普通!")
		case 6:
			fmt.Println("学生等级为及格!")
		default:
			fmt.Println("学生等级为垃圾!")
	}
}
