package main

import(
	"fmt"
	"github.com/wulin05/mashibing_golang/project01/unit2/demo17/test"
)

func main() {
	//如果util.go中定义的是stuNo的话，那么在test.go中是访问不到的，会报错：
	fmt.Println(test.StuNo)
}