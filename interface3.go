package main

import "fmt"

type User struct {
	Id   int
	Name string
}

type Tester interface {
	Test()
}

func (this User) Test() {}

func main() {
	u := User{1, "Tom"}
	i := Tester(u)

	u.Id = 100
	u.Name = "Jack"
	fmt.Println(u, i)

	//.\interface3.go:24: cannot assign to i.(User).Name
	i.(User).Name = "Jack"
	//.\interface3.go:25: cannot take the address of i.(User)
	fmt.Println(&(i.(User)))

}
