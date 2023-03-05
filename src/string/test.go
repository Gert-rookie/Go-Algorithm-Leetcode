package main

import "fmt"

func main() {

	ch := make(chan int, 1)
	x := <-ch
	fmt.Println(x)
	go func() {
		ch <- 1
	}()

}
