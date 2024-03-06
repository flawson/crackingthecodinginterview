package main

import "github.com/flawson/crackingthecodinginterview/chapter2/linkedlist"

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
