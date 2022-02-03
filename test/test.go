package main

import "fmt"

func main() {
	var a = []int{12, 20, 5, 16, 15, 1, 30, 45, 23, 9}
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Print(a)
}
