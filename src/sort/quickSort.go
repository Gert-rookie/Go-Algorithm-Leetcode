package main

import "fmt"

/**
  快速排序
*/
func main() {
	var a = []int{12, 20, 5, 16, 15, 1, 30, 45, 23, 9}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quickSort(a []int, left, right int) {
	if left > right {
		return
	}
	pivot := partition(a, left, right)
	quickSort(a, 0, pivot-1)
	quickSort(a, pivot+1, right)
}

func partition(a []int, left, right int) int {
	pivot := left
	index := left + 1
	for i := index; i <= right; i++ {
		if a[i] < a[pivot] {
			a[i], a[index] = a[index], a[i]
			index++
		}
	}
	a[pivot], a[index-1] = a[index-1], a[pivot]
	return index - 1
}
