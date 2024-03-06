package main

import (
	"fmt"
	"math/rand"

	"github.com/flawson/crackingthecodinginterview/chapter2/linkedlist"
)

func RemoveDups(list *linkedlist.SingleLinkedList) {
	for current := list.Head; current.Next != nil; current = current.Next {
		for marker := current.Next; marker != nil; marker = marker.Next {
			if current.Value == marker.Value {
				tmp := marker.Next
				current.Next = tmp
			}
		}
	}
}

func main() {
	list := linkedlist.NewSingleLinkedList()
	for i := 0; i < 100; i++ {
		list.Insert(linkedlist.NewSingleNode(rand.Intn(100)))
	}
	RemoveDups(list)
	for current := list.Head; current != nil; current = current.Next {
		fmt.Printf("%d\n", current.Value)
	}
}
