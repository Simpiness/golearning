package main

import "fmt"
import "time"

func main() {
	c := make(chan int, 10)
	o := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("recv:", <-c)
		o <- true
	}()

	c <- 100
	fmt.Println("send over ...")
	<-o
}
