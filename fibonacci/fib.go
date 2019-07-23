package main

import "fmt"

func main() {
	res := fibRecurse(20)
	fmt.Print(res)
}

func fibRecurse(n int) int {
	if n <= 1 {
		return n
	}

	return fibRecurse(n-1) + fibRecurse(n-2)
}

func fibLoop(n int) int {
	return 0
}
