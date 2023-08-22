package main

import (
	"errors"
	"fmt"
)

//demo02这个代码与demo01是本质的区别,demo01的通过defer+recover机制来处理错误,避免后续代码不能继续运行
//而这个demo02是通过errors包的New函数来自定义错误,
//其实demo01和demo02的方式,都有一种:我把错误接收了,你程序只管继续执行,不能因为这个错误,就不继续执行了！
//

func main() {
	err := test()
	fmt.Printf("err的类型是: %T \n", err)
	if err != nil {
		fmt.Println("自定义错误:", err) //这种情况程序还会执行下面的代码
		//如果觉得没必要执行,想让程序中的,退出程序,就要借助: builtin包下内置函数: panic
		panic(err) // 这个代码写了就会终止程序,下面的Println语句就不再执行了！
	}
	fmt.Println("上面的test()函数执行成功.....")
	fmt.Println("正常执行下面的逻辑.....")

}

func test() (err error) { //注意,不需要传入参数,但是有返回值: 返回值变量名为err,类型为error
	num1 := 10
	num2 := 2
	if num2 == 0 {
		//抛出自定义错误
		return errors.New("除数不能为0哦~~")
	} else { //如果除数不为0,那么正常执行就可以了
		result := num1 / num2
		fmt.Println(result)
		//如果没有错误,就返回零值
		return nil
	}
}
