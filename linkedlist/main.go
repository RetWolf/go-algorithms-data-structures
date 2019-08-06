package main

import "fmt"

func main() {
	list := SLinkedList{}
	list.addToHead(8)
	list.addToHead(7)
	list.addToTail(7)
	list.addToTail(20)
	list.addToHead(8)
	list.addToHead(8)
	list.addToTail(9)
	list.addToTail(20)
	list.printAll()
	fmt.Println(" ----- ")
	list.removeDuplicates()
	list.printAll()
}
