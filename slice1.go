package main

import "fmt"

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 []int = x[1:3]
	fmt.Println(s1)

	s2 := x[4:]
	fmt.Println(s2)

	s3 := x[:6]
	fmt.Println(s3)

	s4 := x[:]
	fmt.Println(s4)

	//对slice修改就是对底层对象array的修改
	s4[1] = 100
	fmt.Println(x)

	//可以直接创建slice
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	test(s5)
	fmt.Println(s5)

	//new为清零内存，make才是初始化
	s6 := make([]int, 10, 10)
	fmt.Println(s6)
	fmt.Println(len(s6))
	fmt.Println(cap(s6))
}

func test(s []int) {
	s[0] = 100
}
