package main

import "fmt"

func main() {
	s1 := make([]int, 3, 6)
	//添加数据，未超出底层数组容量限制
	s2 := append(s1, 1, 2, 3)
	//append不会调整原slice属性
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	//追加的数据未超出底层容量限制，通过调整s1，依然使用的原数组
	s1 = s1[:cap(s1)]
	fmt.Println(s1, len(s1), cap(s1))

	//再次追加数据（使用了变参）
	//原底层数组已经无法容纳新的数据，将重新分配内存，并拷贝原有数据
	s3 := append(s2, []int{4, 5, 6}...)
	s3[0] = 100
	fmt.Println(s3, len(s3), cap(s3))
	//原slice对象依然是指向旧的底层数组对象
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s1), cap(s1))
	//同步比较s3[0]和s1[0]就可知道是否为原始数组
}
