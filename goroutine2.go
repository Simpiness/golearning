package main

import "fmt"

func test(c chan bool, n int) {
	x := 0
	for i := 0; i < 100000000; i++ {
		x += 1
	}
	fmt.Println(n, x)
	if n == 9 {
		c <- true
	}
}

func main() {
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go test(c, i)
	}
	<-c
}
