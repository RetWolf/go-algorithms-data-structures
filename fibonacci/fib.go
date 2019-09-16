package main

import "fmt"

func main() {
	recurse := fibRecurse(20)
	loop := fibLoop(20)
	fmt.Println(recurse)
	fmt.Println(loop)
}

//  O(2 ^n)
// This function can be significantly improved using memoization, not sure off the top of my head what the big O complexity is for the memoized solution
func fibRecurse(n int) int {
	if n <= 1 {
		return n
	}

	return fibRecurse(n-1) + fibRecurse(n-2)
}

// O(n)
// This is fairly optimal
func fibLoop(n int) int {
	curr := 1
	for prev, i := 0, 1; i < n; i++ {
		prev, curr = curr, curr+prev
	}

	return curr
}
