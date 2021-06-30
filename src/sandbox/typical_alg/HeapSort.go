package main

import (
	"fmt"
)

//O(n*log n)
func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	HeapSort(ar)
	fmt.Println(ar)
}

func HeapSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	heapIt(ar)

	ar[0], ar[len(ar)-1] = ar[len(ar)-1], ar[0]

	HeapSort(ar[:len(ar)-1])
}

func heapIt(ar []int) {

	if len(ar) < 2 {
		return
	}

	if len(ar) == 2 {
		if ar[0] < ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
		return
	}

	if len(ar) > 3 {
		heapIt(ar[1:])
		heapIt(ar[2:])
	}

	if ar[1] > ar[2] {
		if ar[0] < ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
	} else {
		if ar[0] < ar[2] {
			ar[0], ar[2] = ar[2], ar[0]
		}
	}
}
