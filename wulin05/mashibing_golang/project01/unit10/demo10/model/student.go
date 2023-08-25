package model

type Student struct { //这里的关键点是 Student,首字母一定要大写,可以在其它包下访问!!!
	Name string
	Age  int
}

type teacher struct {
	Name string
	Sex  string
	Age  int
}

//工厂模式 方法
func NewTeacher(n string, s string, a int) *teacher {
	return &teacher{n, s, a}
}
