package main

import "fmt"
import "bytes"

func main() {
	sb := bytes.Buffer{}
	sb.WriteString("Hello ")
	sb.WriteString("World ")
	sb.WriteString("!")
	fmt.Println(sb.String())

	s := "abcdefg"
	sub := s[1:3]
	fmt.Printf("%T = %v\n", sub, sub)

	s = "abc"
	fmt.Printf("%c,%02x\n", s[1], s[1])

	bs := []byte(s)
	bs[1] = 'B'
	fmt.Println(string(bs))

	fmt.Println("--------------------------------------")
	s = "a:中国人"
	fmt.Println("utf-8 string", s, len(s))
	/*---string to bytes---*/
	//转化为bytes
	bs = []byte(s)
	bs[0] = 'A'

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%02x ", bs[i])
	}
	//将bytes转化为string
	fmt.Println(string(bs))

	//转化成unicode
	u := []rune(s)
	fmt.Println("unicode len = ", len(u))
	for i := 0; i < len(u); i++ {
		fmt.Printf("%04x ", u[i])
	}
	u[4] = '龙'
	fmt.Println(string(u))

	s = "中国人"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	for i, c := range s {
		fmt.Printf("%d = %c\n", i, c)
	}

}
