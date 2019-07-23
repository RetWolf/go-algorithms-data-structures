package main

import "fmt"

func main() {
	recurse := fibRecurse(20)
	loop := fibLoop(20)
	fmt.Println(recurse)
	fmt.Println(loop)
}

func fibRecurse(n int) int {
	if n <= 1 {
		return n
	}

	return fibRecurse(n-1) + fibRecurse(n-2)
}

func fibLoop(n int) int {
	prev := 0
	current := 1
	for i := 1; i < n; i++ {
		prev, current = current, current+prev
	}

	return current
}
