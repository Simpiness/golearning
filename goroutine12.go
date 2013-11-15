package main

import "fmt"

type Data struct {
	args []int
	ch   chan string
}

func server() chan *Data {
	reqs := make(chan *Data)
	go func() {
		for d := range reqs {
			go serverProcess(d)
		}
	}()
	return reqs
}

func serverProcess(data *Data) {
	x := 0
	for _, i := range data.args {
		x += i
	}
	s := fmt.Sprintf("server : %d", x)
	data.ch <- s
}

func main() {
	/*启动服务器*/
	serverReqs := server()

	/*客户端向服务器发送请求数据，并等待返回结果*/
	data := &Data{[]int{1, 2, 3}, make(chan string)}
	//发送请求数据到服务器
	serverReqs <- data
	//获取服务器返回结果
	fmt.Println(<-data.ch)

	/*关闭服务器*/
	close(serverReqs)
}
