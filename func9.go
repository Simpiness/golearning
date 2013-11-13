package main

import "fmt"

//defer参数定义时值拷贝，可以考虑用指针或者闭包代替
func main() {
	//test()
	x := 10
	defer func(a int) {
		fmt.Println("a = ", a)
	}(x)

	defer fmt.Println("print = ", x)
	x += 100
	fmt.Println(x)
	test()
}

func test() {
	x := 20
	p := &x
	defer func(a int) {
		fmt.Println("x = ", *p)
		fmt.Println("x = ", a)
	}(x)
	x += 100
	*p += 1
	//fmt.Printf("%p = %v\n", p, x)
}
