package main

import "fmt"

/**
  冒泡排序
*/
func main() {
	var a = []int{12, 20, 5, 16, 15, 1, 30, 45, 23, 9}
	bubblingSort(a)
	fmt.Println(a)
}

/**
  冒泡排序
*/
func bubblingSort(a []int) {
	if a == nil || len(a) < 2 {
		return
	}
	// 比较的轮次
	for i := 0; i < len(a)-1; i++ {
		// 两两比较
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
