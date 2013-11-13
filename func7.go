package main

import "fmt"

//闭包指向的是同一个对象x，而不是复制变量x
func test(x int) (f1 func(), f2 func()) {
	fmt.Printf("test : %p = %v\n", &x, x)
	f1 = func() {
		x += 100
		fmt.Printf("f1:%p = %v\n", &x, x)
	}

	f2 = func() {
		fmt.Printf("f2:%p = %v\n", &x, x)
	}

	return
}

func main() {
	f1, f2 := test(100)
	f1()
	f2()
}
