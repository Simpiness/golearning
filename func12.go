package main

import "fmt"
import "errors"

func main() {
	fmt.Println("Hello, World!")
	func() {
		defer func() {
			fmt.Println(recover())
		}()
		panic(errors.New("Error!"))
	}()
}
