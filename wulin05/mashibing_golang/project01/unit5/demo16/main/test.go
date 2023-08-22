package main

import (
	"fmt"
	"time"
)

/*
关于时间和日期的函数,需要导入time包
*/

func main() {
	//获取当前时间,需要调用time包里的Now函数
	now := time.Now() // Now()返回值是一个结构体,类型是: time.Time
	fmt.Printf("当前时间是: %v \n对应的类型是: %T \n", now, now)

	//既然Now()返回值是一个结构体,那么可以调用结构体中的方法:
	fmt.Printf("年: %v \n", now.Year())
	fmt.Printf("月: %v \n", now.Month())      // Auguest
	fmt.Printf("月: %v \n", int(now.Month())) // 8
	fmt.Printf("日: %v \n", now.Day())
	fmt.Printf("时: %v \n", now.Hour())
	fmt.Printf("分: %v \n", now.Minute())
	fmt.Printf("秒: %v \n", now.Second())

	fmt.Println("--------------------------------------------------------")

	//Printf将字符串直接输出:
	fmt.Printf("当前年月日: %d-%d-%d 时分秒: %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//Sprintf可以得到这个字符串,以便后续使用:
	datestr := fmt.Sprintf("当前年月日: %d-%d-%d 时分秒: %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(datestr)

	fmt.Println("--------------------------------------------------------")

	//这个参数字符串"2006/01/02 15:04:05"的各个数字必须是固定的,组合格式可以自己选。
	datestr02 := now.Format("2006/01/02 15:04:05")
	//datestr02 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr02)

	datestr03 := now.Format("2006 15:04") //只留年份、小时和分钟
	fmt.Println(datestr03)

	datestr04 := now.Format("01/02 15:04:05") //只留月份、日期、小时和分钟
	fmt.Println(datestr04)

}
