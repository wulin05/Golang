package main

import (
	"fmt"
)

/*
5. 切片可以动态增长
   append(切片名,data1,data2,...)

其实,append不是直接对slice切片或者切片对应的数组直接追加,而是在底层创建了一个新的数组
将原来alice切片对应的数据拷贝到新数组,再将append追加的数据拷贝到新数组的后面
然后用新的切片slice2底层数组的指向 指向的是新的数组
但是,往往我们在使用追加的时候其实想要做的效果给slice追加:
*/

func main() {
	//定义一个数组
	var intarr [6]int = [6]int{1, 4, 7, 3, 6, 9}
	//定义切片
	var slice []int = intarr[1:4]
	fmt.Println(slice)      // [4 7 3]
	fmt.Println(len(slice)) // 3

	fmt.Println("-------------------------------------------------------------")

	/*

	    //从最上面的注释,可以得知append操作是底层新建了一个数组,用slice2去接收！
	    //那么当然也可以直接让slice自己本身去接收,这样的话,intarr数组还会受slice控制么？
	   slice2 := append(slice,55,77)
	   fmt.Println(slice2) // [4 7 3 55 77]
	    //  slice = append(slice,55,77)
	    //  fmt.Println(slice) // [4 7 3 55 77]

	    fmt.Println("-------------------------------------------------------------")
	    //当经过上一步slice = append(slice,55,77)后,查看intarr数组会不会受slice切片操作的影响
	    fmt.Println("intarr数组是:", intarr) // [1 4 7 3 55 77]
	    slice2[4] = 100
	    fmt.Println(slice2) // [4 7 3 55 100]
	    fmt.Println(intarr) // [1 4 7 3 55 100]

	*/

	//注意下面跟上面从slice2 := append(slice,55,77)到fmt.Println(intarr)的输出区别
	slice = append(slice, 111, 222, 333)
	fmt.Println(slice)  // [4 7 3 111 222 333]
	fmt.Println(intarr) // [1 4 7 3 6 9]

	slice[5] = 444
	fmt.Println(slice)  // [4 7 3 111 222 444]
	fmt.Println(intarr) // [1 4 7 3 6 9]
	//从上面的输出,当slice切片添加内容超出原数组intarr的时候,好像slice就跟原intarr数组
	//没有关系了!!! 对的，如果切片超出老数组的容量,那么就会创建新的数组,导致切片和原数组分离！

	fmt.Println("-------------------------------------------------------------")
	//定义一个新切片,追加到slice切片中去
	slice3 := []int{1000, 2000}
	slice = append(slice, slice3...) // slice3后面的...不能省略,代表追加的slice3是切片
	fmt.Println(slice)               // [4 7 3 111 222 444 1000 2000]
	fmt.Println(intarr)              // [1 4 7 3 6 9],从这可以看出自从slice切片超过intarr数组的长度后,slice切片就跟
	// 原数组intarr没关系了,对slice的操作不会影响到intarr数组了！

}
