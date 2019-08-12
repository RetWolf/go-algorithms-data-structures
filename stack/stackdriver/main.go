package main

import (
	"fmt"

	"github.com/retwolf/go-algorithms-data-structures/stack"
)

func main() {
	as := stack.ArrayStack{}
	as.Push(10)
	as.Push(50)
	fmt.Println(as)
	fmt.Println(as.Peek())
	fmt.Println(as.Pop())
	as.Push(100)
	fmt.Println(as)
	fmt.Println(as.Size())
	fmt.Println(as.Find(120))
	as.Clear()
	fmt.Println(as)
	as.Push(30)
	fmt.Println(as)
}
