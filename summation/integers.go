package main

import "fmt"

func main() {
	sumPrint(1000)
}

func sumPrint(max int) {
	max = max - 1 // If we include 0, we can get rid of this otherwise
	sum := max * ((1 + max) / 2)
	fmt.Print(sum)
}

func sumLoop(max int) int {
	sum := 0
	for i := 0; i < max; i++ {
		sum += i
	}

	return sum
}

func sumFormula(max int) int {
	max = max - 1 // We need this because we include 0
	sum := max * ((1 + max) / 2)

	return sum
}
