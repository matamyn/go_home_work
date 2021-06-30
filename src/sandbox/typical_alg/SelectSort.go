package main

import (
	"fmt"
)

//O(n*n)
func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	SelectSort(ar)
	fmt.Println(ar)
}

func SelectSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[min] > ar[j] {
				min = j
			}
		}

		if min != i {
			ar[i], ar[min] = ar[min], ar[i]
		}
	}
}
