package main

import (
	"fmt"
	"github.com/wulin05/mashibing_golang/project01/unit10/demo11/model"
)

/*
封装的引入
1. 封装(encapsulation)就是把抽象出的字段和对字段的操作封装在一起,数据被包含在内部,程序的其他包只有通过被授权的操作方法,才能对字段进行操作。
2. 封装的好处:
   (1).隐藏实现细节
   (2).可以对数据进行验证,保证安全合理

如何实现封装：
1. 建议将结构体、字段(属性)的首字母小写(其他包不能使用,类似private,实际开发不小写也可以,因为封装没有那么严格)
2. 给结构体所在的包提供一个工厂模式的函数,首字母大写(类似一个构造函数)
3. 提供一个首字母大写的Set方法(类似与其他语言的public),用于对属性判断并赋值
   func (var 结构体类型名) SetXxx(参数列表){
	   //加入数据验证的业务逻辑
	   var.Age = 参数
   }
4. 提供一个首字母大写的Get方法(类似其他语言的public),用于获取属性的值
   func (var 结构体类型名) GetXxx()(返回值列表){
	   return var.字段
   }

*/

//代码实现：

func main() {

	//创建一个person的结构体实例,应该这么理解，NewPerson是函数,返回值是一个person类型的结构体,这个person结构体是这样的：{Name string, age int}
	p := model.NewPerson("丽丽")
	p.SetAge(30) //由于p这个变量类型是person类型的结构体,所以,可以调用person结构体的方法SetAge(),就是给p结构体的age赋值
	// 所以经过上面的两个步骤,person结构体实例p的Name和age都有值了,下面就可以打印输出！
	fmt.Println(p.Name)     // 丽丽
	fmt.Println(p.GetAge()) // 30
	fmt.Println(*p)         // {丽丽 30}

}
