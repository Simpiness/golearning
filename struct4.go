package main

import "fmt"

type User struct {
	Id   int
	Name string
}

type Student struct {
}

type Manager struct {
	Group string
	User
	Student
}

func (this *Student) Test() {
	fmt.Printf("Test Student address = %p\n", this)
}

func (this *User) Test() {
	fmt.Printf("Test User address = %p\n", this)
}

func main() {
	m := &Manager{User: User{1, "Jack"}, Student: Student{}, Group: "IT"}
	fmt.Printf("m.address = %p\n", m)
	fmt.Printf("m.User.address = %p\n", &m.User)
	fmt.Printf("m.Student.address = %p\n", &m.Student)
	//.\struct4.go:32: ambiguous selector m.Test
	//m.Test()
	m.User.Test()
	m.Student.Test()
}
