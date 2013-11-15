package main

import "fmt"

type User struct {
	Id   int
	Name string
}

type Tester interface {
	Test()
}

type Stringer interface {
	String()
}

func (this User) Test() {
	fmt.Println("Test:", this)
}

func (this *User) String() {
	fmt.Printf("String Id = %d,Name = %s\n", this.Id, this.Name)
}

//T 仅拥有属于 T 类型的⽅方法集，⽽而 *T 则同时拥有 T + *T ⽅方法集
//基于 T 实现⽅方法，表⽰示同时实现了 interface(T) 和 interface(*T) 接⼝口。
//基于 *T 实现⽅方法，那就只能是对 interface(*T) 实现接⼝口。
func main() {
	u := User{1, "Tom"}
	p := &u

	//User.Test()同时实现了Tester(User)和Tester(*User)接口
	Tester(u).Test()
	Tester(p).Test()

	//.\interface4.go:34: cannot convert u (type User) to type Stringer:
	//  User does not implement Stringer (String method requires pointer receiver)
	//Stringer(u).String()
	Stringer(p).String()
}
