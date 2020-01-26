package main

import (
	"fmt"
	"time"
)

// 定义goroutine 1
func Echo(out chan<- string) {   // 定义输出通道类型
	time.Sleep(1*time.Second)
	out <- "Hello World"

}

// 定义goroutine 2
func Receive(out chan string, in chan string) { // 定义输出通道类型和输入类型
	temp := <-in // 阻塞等待echo的通道的返回
	out <- temp

}


func main() {
	echo := make(chan string)
	receive := make(chan string)

	go Echo(echo)
	go Receive(receive, echo)

	getStr := <-receive   // 接收goroutine 2的返回

	fmt.Println(getStr)
}
