package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {
	a := [...]User{
		{0, "User0"},
		{1, "User1"},
	}

	b := [...]*User{
		&User{0, "User0"},
		&User{1, "User1"},
	}

	c := [...]*User{
		{0, "User0"},
		{1, "User1"},
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
