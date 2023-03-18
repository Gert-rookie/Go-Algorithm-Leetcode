package main

import "fmt"

func main() {
	// quick sort
	var a = []int{12, 20, 5, 16, 15, 1, 30, 45, 23, 9}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quickSort(a []int, left int, right int) {
	if left > right {
		return
	}
	p := partition(a, left, right)
	quickSort(a, 0, p-1)
	quickSort(a, p+1, right)

}

func partition(a []int, left int, right int) int {
	p := left
	index := left + 1
	for i := index; i <= right; i++ {
		if a[p] > a[i] {
			a[i], a[index] = a[index], a[i]
			index++
		}
	}
	a[p], a[index-1] = a[index-1], a[p]
	return index - 1
}
