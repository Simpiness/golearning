package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {
	users := map[string]User{
		"a": User{1, "User1"},
	}
	fmt.Println(users)
	//map[key]返回的只是一个“临时值”拷贝，修改自身状态没有任何意义
	//只能重新value赋值或者改用指针修改所引用内存
	//cannot assign to users["a"].Name
	//users["a"].Name = "Jack"

	u := users["a"]
	u.Name = "Jack"
	//重新value赋值
	users["a"] = u
	fmt.Println(users)
	//使用指针类型
	users2 := map[string]*User{
		"a": &User{1, "User1"},
	}
	fmt.Println(users2["a"])
	users2["a"].Name = "Tom"
	fmt.Println(users2["a"])
}
