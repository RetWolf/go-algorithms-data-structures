package main

import "fmt"

// SLinkedList contains a head and a tail, both of which are pointers to SLLNodes
type SLinkedList struct {
	head *SLLNode
	tail *SLLNode
}

func (list *SLinkedList) isEmpty() bool {
	return list.head == nil
}

func (list *SLinkedList) addToHead(el int) {
	list.head = &SLLNode{
		val:  el,
		next: list.head,
	}

	if list.tail == nil {
		list.tail = list.head
	}
}

func (list *SLinkedList) addToTail(el int) {
	if !list.isEmpty() {
		list.tail.next = &SLLNode{
			val: el,
		}
		list.tail = list.tail.next
	} else {
		list.head = &SLLNode{
			val: el,
		}
		list.tail = list.head
	}
}

func (list *SLinkedList) deleteFromHead() int {
	el := list.head.val
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.head = list.head.next
	}

	return el
}

func (list *SLinkedList) deleteFromTail() int {
	el := list.tail.val
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		temp := list.head
		for temp.next != list.tail {
			temp = temp.next
		}

		list.tail = temp
		list.tail.next = nil
	}

	return el
}

func (list *SLinkedList) printAll() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}

func (list *SLinkedList) isInList(el int) bool {
	temp := list.head
	for temp != nil && temp.val != el {
		temp = temp.next
	}

	return temp != nil
}

func (list *SLinkedList) delete(el int) {
	if !list.isEmpty() {
		if list.head == list.tail && el == list.head.val {
			list.head = nil
			list.tail = nil
		} else if el == list.head.val {
			list.head = list.head.next
		} else {
			prev := list.head
			lead := prev.next
			for lead != nil && lead.val != el {
				prev = prev.next
				lead = lead.next
			}

			if lead != nil {
				prev.next = lead.next
				if lead == list.tail {
					list.tail = prev
				}
			}
		}
	}
}
