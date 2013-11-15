package main

import (
	"fmt"
	"sync"
	"time"
)

var sem = make(chan int, 2)
var wg = sync.WaitGroup{}

func worker(i int) {
	sem <- 1
	fmt.Println(time.Now().Format("04:05"), i)
	time.Sleep(1 * time.Second)
	<-sem
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}
