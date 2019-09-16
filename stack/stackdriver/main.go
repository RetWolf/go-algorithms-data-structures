package main

import (
	"fmt"

	"github.com/retwolf/go-algorithms-data-structures/stack"
)

func main() {
	as := stack.ArrayStack{}
	for i := 0; i < 500; i++ {
		as.Push(10)
	}

	fmt.Println(as.Size())
	fmt.Println(as.Pop())
}
