package main
//首先是无缓冲channel,用于保证并发的顺序执行.




import (
	"fmt"
"time"
)
var done = make(chan string)  // 创建一个channel
func Hello() {
	fmt.Println("Hello")
	time.Sleep(1*time.Second)
	done <- "World"
}
func main1() {

	go Hello()
	go fmt.Println(<-done)
	go fmt.Println("112121")
	 fmt.Println("final")
}
/*
Hello
World
*/
