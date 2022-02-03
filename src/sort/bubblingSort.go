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
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ { // 两两排序
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
