package main

import "fmt"
import "runtime"

func test() {
	fmt.Println(runtime.Caller(1))
}

func main() {
	test()
}
