package main

import "fmt"

func main() {
	a := 0x1234
	b := 1234
	c := 256
	fmt.Printf("%x\n", uint8(a))
	fmt.Printf("%d\n", int(b))
	fmt.Printf("%f\n", float64(c))
}

// *Point(b)  		//same as *(Point(p))
// (*Point)(p) 		//p is converted to (*Point)
// <-chan int(c) 	//same as <-(chan int(c))
// (<-chan int)(c)	//c is converted to (<-chan int)
