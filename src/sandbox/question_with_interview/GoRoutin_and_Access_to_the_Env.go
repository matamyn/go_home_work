package main

import (
	"fmt"
	"sync"
)

func TestGoRout(i int) func() int {
	tmp := i
	return func() int {
		tmp = tmp + 1
		return tmp
	}
}

func main() {
	var wg sync.WaitGroup

	tmp := TestGoRout(5)
	fmt.Println(tmp())

	wg.Add(1)
	go func() {
		fmt.Println(tmp())
		wg.Done()
	}()
	go fmt.Println(tmp())

	wg.Wait()
	fmt.Println(tmp())
}
