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

func (list *SLinkedList) sum() int {
	if list.isEmpty() {
		return 0
	}

	total := 0
	for temp := list.head; temp != nil; temp = temp.next {
		total += temp.val
	}
	return total
}

func (list *SLinkedList) max() int {
	if list.isEmpty() {
		return 0
	}

	max := list.head.val
	for temp := list.head; temp != nil; temp = temp.next {
		if temp.val > max {
			max = temp.val
		}
	}
	return max
}

func (list *SLinkedList) min() int {
	if list.isEmpty() {
		return 0
	}

	min := list.head.val
	for temp := list.head; temp != nil; temp = temp.next {
		if temp.val < min {
			min = temp.val
		}
	}
	return min
}

func (list *SLinkedList) removeDuplicates() {
	values := make(map[string]int)
	values[fmt.Sprintf("%v", list.head.val)] = 1

	prev, curr := list.head, list.head.next
	for curr != nil {
		if values[fmt.Sprintf("%v", curr.val)] == 1 {
			// Value already exists in list, need to remove
			prev.next = curr.next
			curr = curr.next
		} else {
			// Value hasn't been found in list yet, increment map
			values[fmt.Sprintf("%v", curr.val)] = 1
			prev, curr = curr, curr.next
		}
	}
}
