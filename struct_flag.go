package main

import "fmt"
import "reflect"

type User struct {
	Name string "姓名"
	Age  int    "年龄"
}

func main() {
	u := User{"Jack", 26}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s (%s = %v)\n", f.Tag, f.Name, v.Field(i).Interface())
	}
}
