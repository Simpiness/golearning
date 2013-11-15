package main

import "time"
import "fmt"

func test(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("go...")
	c <- 1
}

func main() {
	c := make(chan int)
	go test(c)
	fmt.Println("hi")
	<-c
	fmt.Println("over!")
}
