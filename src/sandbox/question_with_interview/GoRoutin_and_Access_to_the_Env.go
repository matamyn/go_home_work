package main

import (
	"fmt"
	"sync"
	"time"
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
	//Функции в горутинах
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

	wg.Add(2) // в группе две горутины
	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}

	// вызываем горутины
	go work(1)
	go work(2)

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
}
