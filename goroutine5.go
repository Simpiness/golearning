package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		//如果不关闭，引发错误
		//fatal error: all goroutines are asleep - deadlock!
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
