package main

import "fmt"

/*
只有发送端 (另⼀一端正在等待接收) 才能 close，只有接收端才能获得关闭状态。close 调用不是必
须的，但如果接收端使⽤用迭代器或者循环接收数据，则必须调⽤用，否则可能导致接收端 throw: all
goroutines are asleep - deadlock!
*/
func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		//fatal error: all goroutines are asleep - deadlock!
		close(c)
	}()

	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
