package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriterTest struct {
}

func (rw *ReadWriterTest) Read() {
	fmt.Println("Read")
}

func (rw *ReadWriterTest) Write() {
	fmt.Println("Write")
}

func main() {
	t := ReadWriterTest{}
	var rw ReadWriter = &t
	rw.Read()
	rw.Write()
}
