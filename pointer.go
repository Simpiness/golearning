package main

import "fmt"
import "unsafe"

type User struct {
	Id   int
	Name string
}

const N int = int(unsafe.Sizeof(0))

func main() {
	i := 100
	var p *int = &i
	fmt.Println(*p)

	up := &User{Id: 1, Name: "jack"}
	fmt.Println(up)

	u2 := *up
	u2.Name = "Tom"
	fmt.Println(up, u2)
	fmt.Println("----------------------------")
	x := 0x1234
	p1 := unsafe.Pointer(&x)
	p2 := (*[N]int)(p1)
	for i, m := 0, len(p2); i < m; i++ {
		fmt.Printf("%02X ", p2[i])
	}
}
