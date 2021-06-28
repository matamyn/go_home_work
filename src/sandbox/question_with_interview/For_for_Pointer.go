package main

import (
	"fmt"
)

//Логика памяти Циклов!

func test(tmp []*int) []*int {
	var sum = 0
	for i, t := range tmp {
		fmt.Printf("%d ", t)
		var x = *t * 2
		tmp[i] = &x
		sum += x
	}
	return append(tmp, &sum)
}

func main() {
	var numbers []*int
	var arr = []int{1, 2, 3, 4}
	for _, value := range arr {
		fmt.Printf("%d ", &value)
		var tmp int = value
		numbers = append(numbers, &tmp)
	}

	tmp := test(numbers)

	for _, number := range tmp {
		fmt.Printf("%d ", *number)
	}

}
