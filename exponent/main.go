package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(powerIterative(2, 5))
}

func powerRecurse(base, expo int) int {
	if expo == 0 {
		return 1
	}

	return base * powerRecurse(base, expo-1)
}

func powerIterative(base, expo int) int {
	res := 1
	for expo > 0 {
		if expo&1 == 1 {
			res = res * base
		}
		expo = expo >> 1
		base = base * base
	}

	return res
}

func power(base, expo float64) float64 {
	return math.Pow(base, expo)
}
