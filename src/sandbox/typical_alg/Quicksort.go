package main

import (
	"fmt"
)

//max - O(n*n)
//min - O(log n)
func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	Quicksort(ar)
	fmt.Println(ar)
}

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}
		ar[left], ar[right] = ar[right], ar[left]

	}
}
