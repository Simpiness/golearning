package main

import "fmt"

type User struct {
	Name string
}

func (u User) Test() {
	fmt.Println(u)
}

func main() {
	u := User{"Tom"}
	//这里没有发生函数调用
	f := u.Test

	u.Name = "xxx"
	fmt.Println(u)

	f()
}
