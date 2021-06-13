package main

import (
	"container/list"
	"fmt"
)

// ref: https://golang.org/pkg/container/list/

func main() {
	// initialized list
	l := list.New()
	// inserts a new element
	l.PushBack(1) // 1

	m := list.New()
	m.PushBack(2)
	m.PushBack(3)
	fmt.Printf("number of elements of list m: %d\n", m.Len())

	// inserts a copy of another list at the back of list l
	l.PushBackList(m) // 1 2 3
	fmt.Println("=== after PushBackList(m) ===")
	// To iterate over a list
	printList(l)

	// inserts a new element e with value v at the front of list l
	l.PushFront(5) // 5, 1, 2, 3

	// inserts a copy of another list at the front of list l.
	l.PushFrontList(m) // 2, 3, 5, 1 , 2, 3
	fmt.Println("=== after PushFront, PushFrontList(m) ===")
	printList(l)

	// Back returns the last element
	// Remove removes element from l
	l.Remove(l.Back())
	fmt.Println("=== after Remove ===")
	printList(l)

	var target *list.Element

	for e := l.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(int); ok && v == 5 {
			// InsertBefore もある
			target = l.InsertAfter(10, e)
			break
		}
	}
	fmt.Println("=== after  InsertAfter ===")
	printList(l)

	// MoveToBack, MoveBefore などもある
	l.MoveToFront(target)
	fmt.Println("=== after MoveToFront ===")
	printList(l)
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
