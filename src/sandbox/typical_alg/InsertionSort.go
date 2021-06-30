package main

import (
	"fmt"
)

//O(n*n)
func main() {
	ar1 := []int{3, 4, 1, 2, 5, 7, -1, 0}
	InsertionSort(ar1)
	fmt.Println(ar1)
	//Max O(n*n)
	//Arv O(n*logN)
	ar2 := []int{3, 4, 1, 2, 5, 7, -1, 0}
	ShellSort(ar2)
	fmt.Println(ar2)
}

func InsertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}

func ShellSort(ar []int) {
	for gap := len(ar) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(ar); i++ {
			x := ar[i]
			j := i
			for ; j >= gap && ar[j-gap] > x; j -= gap {
				ar[j] = ar[j-gap]
			}
			ar[j] = x
		}
	}
}
