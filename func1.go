package main

import "fmt"

type callback func(s string)

func test(a, b int, sum func(int, int) int) {
	fmt.Println(sum(a, b))
}

func main() {
	var cb callback = func(s string) {
		fmt.Println(s)
	}

	cb("Hello,World!")
	test(1, 2, func(a, b int) int {
		return a + b
	})
}
