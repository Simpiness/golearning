package main

import "fmt"
import "reflect"

/*
	T 方法仅拥有T Receiver方法
	*T 方法集则包含全部方法（T + *T）
*/
type User struct {
	Name string
}

func (this User) TestValue() {
	fmt.Printf("V:%p,%v\n", &this, this)
}

func (this *User) TestPointer() {
	fmt.Printf("P:%p,%v\n", this, *this)
}

func main() {
	u := User{"Tom"}
	fmt.Printf("u:%p,%v\n", &u, u)
	u.TestValue()
	u.TestPointer()

	User.TestValue(u)
	//User.TestPointer undefined (type User has no method TestPointer)
	//User.TestPointer(&u)

	//(*User)方法集中包好了TestValue，TestPointer
	(*User).TestValue(&u)
	(*User).TestPointer(&u)
	//反射查看签名变化
	m, _ := reflect.TypeOf(&u).MethodByName("TestValue")
	fmt.Println(m)

	u1 := &User{"Jack"}
	u1.TestValue()
	u1.TestPointer()

	(*User).TestValue(u1)
	(*User).TestPointer(u1)
}
