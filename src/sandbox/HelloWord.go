package main

import "fmt"

func main() {
	fmt.Println("GO") // ⚡ index out of range [0] with length 0

	array := make([]int, 10, 10)
	fmt.Println(array[9]) // ⚡ index out of range [0] with length 0

	symbols := [...]rune{'a', 'b', 'c', 'd', 'e'}
	fmt.Println(symbols[4]) // ⚡ index out of range [0] with length 0

	var decisions [4]uint     // [0][1][2][3]
	fmt.Println(decisions[3]) // invalid array index 4 🐞
	fmt.Println("Hello, 世界")

	const nihongo = "日本語"
	fmt.Printf("len: %d\n", len(nihongo))
	//fmt.Printf("cap: %d\n", cap(nihongo)) // 🙋 скомпилируется ?
}
