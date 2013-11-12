package main

import "fmt"
import "unsafe"

type ByteSize int64

func main() {
	const (
		a = "abc"
		b
	)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(a, b)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		c = "abc"
		d = unsafe.Sizeof(c)
	)

	fmt.Println(c, d)

	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)
	fmt.Println(GB, GB/1024/1024/1024/1024)

	const (
		A, B = iota, iota
		C, D
		E, F
	)
	fmt.Println(A, B, C, D, E, F)

	const (
		a1 = iota
		b1
		c1
		d1 = "ha"
		e1
		f1 = 100
		g1
		h1 = iota
		i1
	)
	fmt.Println(a1, b1, c1, d1, e1, f1, g1, h1, i1)
}
