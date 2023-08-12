package main

import "fmt"

func main() {
	//功能: 输出1-100中被6整除的数:
	//方式一:
	for i := 1; i <= 100; i++{
		if i % 6 == 0{
			fmt.Println(i)
		}
	}

	fmt.Println("---------------------1")

	//方式二:
	//使用continue
	for j := 1; j <= 100; j++{
		if j % 6 != 0 {
			continue //结束本次循环,继续下一次循环
		}
		fmt.Println(j)
	}

	fmt.Println("---------------------2")

	//举例三:
	for m := 1; m <= 5; m++{
		for n := 2; n <= 4; n++{
			if m == 2 && n == 2 {
				continue //从运行结果看出,只有m:2,n:2没有输出,其他都输出了
			}            //说明: continue作用是结束当次最近的for循环,fmt.Printf()这句语句被跳过了
			fmt.Printf("m: %v, n: %v \n", m, n)
		}
	}

	fmt.Println("---------------------3")

	//举例四:
	//continue也可以使用标签: lable
	lable:
	for m := 1; m <= 5; m++{
		for n := 2; n <= 4; n++{
			if m == 2 && n == 2 {
				continue lable       //查看结果可知: m==2&&n==2的时候,外层循环结束,继续执行m==3的时候的循环!
			}                          //所以,从m==2&&n==2及以后的结果都被跳过了,下次输出是从m==3&&n==2继续开始循环
			fmt.Printf("m: %v, n: %v \n", m, n)    //可以看出continue的这种方式等于break
		}
	}

}
