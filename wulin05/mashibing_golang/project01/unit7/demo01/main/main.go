package main

import (
	"fmt"
)

//数组的引入

func main() {
	//定义一个数组: var 数组名 [数组大小]数据类型 , 如下:
	var scores [5]int
	scores[0] = 99
	scores[1] = 88
	scores[2] = 77
	scores[3] = 66
	scores[4] = 55

	//求和:
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	//求平均数:
	avg := sum / len(scores)

	//输出:
	fmt.Printf("成绩的综合为: %v, 成绩的平均数为: %v", sum, avg)
}
