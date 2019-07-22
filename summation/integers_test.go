package main

import "testing"

func BenchmarkThousand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumFormula(1000)
	}
}

func BenchmarkBillion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumFormula(1000000000)
	}
}

func BenchmarkQuintillion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumFormula(1000000000000000000)
	}
}
