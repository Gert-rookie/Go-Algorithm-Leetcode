package main

import "fmt"

func main() {
	var a = []int{12, 20, 5, 16, 15, 1, 30, 45, 23, 9}
	chooseSort(a)
	fmt.Println(a)
}

/**
  选择排序
  选出最小的数据,放在前面
*/
func chooseSort(a []int) {
	if a == nil || len(a) < 2 {
		return
	}
	for i := 0; i < len(a)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(a); j++ {
			if a[minIndex] > a[j] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}
