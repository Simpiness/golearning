package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func change(a, b int) (x, y int) {
	x = a + 100
	y = b + 100
	return
}

func main() {
	a, b := 1, 2
	a, b = swap(a, b)
	fmt.Println(a, b)

	c, d := change(1, 2)
	fmt.Println(c, d)
}
