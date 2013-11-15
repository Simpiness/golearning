package main

import "fmt"

//channel指定为单向通道，<-chan int仅能接收，
//"chan <- int"仅能发送
func recv(c <-chan int, over chan<- bool) {
	for v := range c {
		fmt.Println(v)
	}

	over <- true
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	//只能发送端关闭
	close(c)
}

func main() {
	c := make(chan int)
	o := make(chan bool)

	go recv(c, o)
	go send(c)

	i := <-o
	fmt.Println(i)
}
