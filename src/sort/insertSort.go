package main

import "fmt"

/**
插入排序
*/

func main() {
	var a = []int{12, 20, 5, 16, 15, 1, 30, 45, 23, 9}
	insertSort(a)
	fmt.Println(a)
}

func insertSort(a []int) {
	if a == nil || len(a) <= 2 {
		return
	}
	// 从第二个数开始，往前比较，若前一个数大则交换，一直比较到第一个数
	// 比较的轮次
	for i := 1; i < len(a); i++ {
		// 两两比较
		for j := i - 1; j >= 0 && a[j] > a[j+1]; j-- {
			a[j], a[j+1] = a[j+1], a[j]
		}
	}

}
