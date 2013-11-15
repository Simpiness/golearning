package main

import "fmt"

type User struct {
	Id   int
	Name string
}

type Manager struct {
	Group string
	User
}

type Tester interface {
	Test()
}

func (this *User) Test() {
	fmt.Println(this)
}

func main() {
	u := User{1, "Tom"}
	m := Manager{User: User{2, "Jack"}, Group: "IT"}
	var i Tester
	//我们为 *User 定义了 Test()，⾃自然表⽰示实现了该接⼝口
	i = &u
	i.Test()

	//匿名字段，同样也拥有 Test()
	i = &m
	i.Test()
}
