/*
 go lang zero values
*/
package main

import "fmt"

func main() {
	var a int
	var b int32
	var c int64
	var d byte
	var e float32
	var f float64
	var g bool
	var h string
	var i *int
	var j func(int) int
	var k interface{}
	var l []int
	var m (chan int)
	var n map[string]int

	fmt.Println("a = 0:", a == 0)
	fmt.Println("b = 0:", b == 0)
	fmt.Println("c = 0:", c == 0)
	fmt.Println("d = 0:", d == 0)
	fmt.Println("e = 0:", e == 0)
	fmt.Println("f = 0:", f == 0)
	fmt.Println("g = false:", g == false)
	fmt.Println("h = \"\":", h == "")
	fmt.Println("i = nil:", i == nil)
	fmt.Println("j = nil:", j == nil)
	fmt.Println("k = nil:", k == nil)
	fmt.Println("l = nil:", l == nil)
	fmt.Println("m = nil:", m == nil)
	fmt.Println("n = nil:", n == nil)
}
