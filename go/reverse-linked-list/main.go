package main

import (
	"fmt"
	"reverse-linked-list/list"
)

func main() {
	entries := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10}
	llist := list.NewEmptyLList[int]()
	fmt.Println("source: ")
	for _, e := range entries {
		if err := llist.Add(e); err != nil {
			panic(err.Error())
		}
	}
	llist.Print(" ")
	fmt.Println()
	invllist := list.ReverseLList[int](llist)
	fmt.Println("result: ")
	invllist.Print(" ")
	fmt.Println()
}
