package main

import "fmt"
import "time"

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("go...")
	}()

	time.Sleep(5 * time.Second)
}
