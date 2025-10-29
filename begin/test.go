package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var ch1 = make(chan string, 10)
	var ch2 = make(chan string, 10)
	go task1(ch1, ch2)
	wg.Add(2)
	go task2(ch1)
	go task2(ch2)
	wg.Wait()
}
func task1(ch1, ch2 chan string) {
	for i := 1; i < 10; i++ {
		//fmt.Println("发送", "a")
		ch1 <- "a"
		ch2 <- "b"
	}
	close(ch1)
	close(ch2)

}

func task2(ch chan string) {
	for val := range ch {
		fmt.Println("接收", val)
	}
	wg.Done()
}
