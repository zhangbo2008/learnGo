package main



/**
 * 并发编程，切片的线程安全性问题
 */


import (
"fmt"
"sync"
"time"
)

var list []int = []int{}
var wgList sync.WaitGroup = sync.WaitGroup{} //利用add ,wait ,done 3个函数来控制.
var muList sync.Mutex = sync.Mutex{}

func main() {
	// 并发启动的协程数量----第一次测试,不安全的
	max := 10000
	fmt.Printf("list add num=%d\n", max)
	wgList.Add(max)   //先吧信号量加到1万.
	time1 := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		go addNotSafe()   //然后每运行完一个信号量就-1
	}
	wgList.Wait()   //wait函数表示都运行完了.
	time2 := time.Now().UnixNano()
	fmt.Printf("list len=%d, time=%f ms\n", len(list), (float64((time2-time1))/1000000))










	// 覆盖后再执行一次          第二次测试安全的,为什么加锁更快了,很费解???????
	list = []int{}
	fmt.Printf("new list add num=%d\n", max)
	wgList.Add(max)
	time3 := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		go addSafe()
	}
	wgList.Wait()
	time4 := time.Now().UnixNano()
	fmt.Printf("new list len=%d, time=%f ms\n", len(list), (float64((time4-time3))/1000000))
}

// 线程不安全的方法
func addNotSafe() {
	list = append(list, 1)
	wgList.Done()       //底层就是信号量-1
}

// 线程安全的方法，增加了互斥锁
func addSafe() {
	muList.Lock()
	list = append(list, 1)
	muList.Unlock()
	wgList.Done()
}