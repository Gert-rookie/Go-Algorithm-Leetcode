package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	rose := <-ch
	ch <- "玫瑰花"
	fmt.Println(rose)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("1 子goroutine执行完毕")
	}()
	go func() {

		fmt.Println("2 子goroutine执行完毕")

	}()
	wg.Wait()
	fmt.Println("主goroutine执行完毕")

}
