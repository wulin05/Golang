package model

import "fmt"

type person struct { //这里的关键点是person,首字母小写,这样其他包就不能访问
	Name string
	age  int //重点：age首字母是小写,其它包不能访问
}

// 定义工厂模式下的函数：NewPerson
func NewPerson(name string) *person { // NewPerson函数首字母是大写N,参数名为name的字符串类型,返回值是person结构体的指针
	return &person{ // 所以这边用return,返回一个person结构体类型的指针：&person
		Name: name,
	}
}

// 定义工厂模式下 person结构体的方法：
// 定义Set和Get方法,对age字段进行封装,因为在方法中可以加一系列的限制操作,确保被封装字段的安全合理性
func (p *person) SetAge(age int) { // 定义一个结构体为 person 的方法：SetAge, 方法SetAge的传入参数是int类型的参数：age,没有返回值,所以(age int)后没有内容了
	if age > 0 && age < 150 { // 这里的 age 是SetAge方法的传入参数：(age int)
		p.age = age
	} else {
		fmt.Println("对不起,你传入的年龄范围不正确")
	}
}

func (p *person) GetAge() int { // 注意这个GetAge方法就没有传入参数了,因为它是获取返回值的,所以在返回值的位置是int,表示接收到的返回值是int类型
	return p.age
}
