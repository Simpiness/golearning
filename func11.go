package main

import "log"

func main() {
	//错误，无法捕获,recover只能捕获defer宿主函数调用堆栈错误
	/*defer func() {
		func() {
			fmt.Println(recover())
		}()
	}()*/

	//错误；recover只能在defer函数内部调用才有效
	//defer recover()
	//defer fmt.Println(recover())

	//正确方式
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	panic("abc")
}
