package main

import (
	"fmt"
)

//数组的输入：Scanln Scanf 在unit3的demo06有详细介绍

func main() {
	var scores [5]int
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第%d个学生的成绩:", i+1)
		/*
			方式一：
			fmt.Scanln(&scores[i])
			方式二：
			fmt.Scanf("%d", &scores[i])
			fmt.Scanln()   // 清除换行符,否则使用Scanf的方式会有问题
		*/
		var input string
		fmt.Scanln(&input)
		fmt.Sscanf(input, "%d", &scores[i])
	}
	//查看录入的成绩：
	fmt.Printf("录入的学生成绩如下: %v \n", scores)

	//求和
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	//求平均值:
	avg := sum / len(scores)

	//打印输出:
	fmt.Printf("成绩总和为: %v, 成绩的平均数为: %v\n", sum, avg)

	//展示班级的每个学生的成绩: 数组进行遍历

	//方式一：普通for循环
	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的成绩为: %d \n", i+1, scores[i])
	}
	fmt.Println("--------------------------------------------------")
	//方式二：for range 在unit4的demo05也有for range的介绍
	/*
		for range结构是Go语言特有的一种迭代结构,可以遍历数组、切片、字符串、map及
		通道, for range语法上类似于其他语言中的foreach语句,一般形式为:
		for key,val := range coll {
			...
		}
		coll:遍历的内容(这里是数组)
		每次遍历得到的索引用key接收
		每次遍历得到的索引位置上的值用val接收
		key、val名字随便起：k/v, key/value, a/b.....
		key/val是局部变量,只在当前循环范围
	*/
	for key, value := range scores {
		fmt.Printf("第%d个学生的成绩为: %d \n", key+1, value)
	}
	fmt.Println("-----------------------------------------------")

	//不想要接收key值,只要输出学生成绩:
	for _, val := range scores {
		fmt.Printf("学生的成绩为: %d\n", val)
	}
}
