package main

import (
	"fmt"
)

//O(n*n)
func main() {
	ar1 := []int{3, 4, 1, 2, 5, 7, -1, 0}
	ar2 := []int{3, 4, 1, 2, 5, 7, -1, 0}
	ar3 := []int{3, 4, 1, 2, 5, 7, -1, 0}
	BubbleSort1(ar1)
	BubbleSort2(ar2)
	BubbleSort3(ar3)
	fmt.Println(ar1)
	fmt.Println(ar2)
	fmt.Println(ar3)

}

//default
func BubbleSort1(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
}

//Пузырь вверх гоним
func BubbleSort2(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := len(ar) - 1; j > i; j-- {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]

			}
		}
	}
}

//sinking sort
func BubbleSort3(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}
