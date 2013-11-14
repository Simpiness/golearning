package main

import "fmt"

type User struct {
	Id   int
	Name string
}

type Manager struct {
	User
	Group string
}

func (this *User) Test() {
	fmt.Println("User Test:", this.Id, this.Name)
}

func (this User) ToString() string {
	return fmt.Sprintf("User ToString:[%d] %s ", this.Id, this.Name)
}

func (this Manager) Test() {
	fmt.Println("Manager Test:", this.Id, this.Name)
}

//实现override效果
func (this Manager) ToString() string {
	return fmt.Sprintf("Manager ToString:[%d] %s", this.Id, this.Name)
}

func main() {
	m := &Manager{User{1, "Jack"}, "IT"}
	//Manager.Test的receiver不是指针
	m.Test()
	//可以直接访问匿名字段（匿名类型或者匿名指针类型）的方法，类似继承
	fmt.Println(m.ToString())
}
