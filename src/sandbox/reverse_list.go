package main

import "fmt"

type query struct {
	val  int
	next *query
}

func createList() *query {

	pHead := &query{1, nil}

	pCurr := pHead
	for i := 0; i <= 10; i++ {
		pItem := &query{i, nil}
		pCurr.next = pItem
		pCurr = pItem
	}

	return pHead
}

func printList(pList *query) {

	pCurr := pList
	for {
		fmt.Printf("%d", pCurr.val)

		if pCurr.next != nil {
			pCurr = pCurr.next
		} else {
			break
		}
	}
	fmt.Println("")
}
func reverse_list(pList *query) *query {
	pCurr := pList

	var pTop *query = nil

	for {
		if pCurr == nil {
			break
		}
		pTemp := pCurr.next
		pCurr.next = pTop
		pTop = pCurr
		pCurr = pTemp
	}

	return pTop
}

func main() {
	pHead := createList()
	printList(pHead)
	printList(reverse_list(pHead))

}
