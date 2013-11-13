package main

import "fmt"

//执行顺序
//x=100 -> call defer func - > ret
func test() (x int) {
	defer func() {
		fmt.Println(x)
	}()
	return 100
}

func main() {
	test()
}
