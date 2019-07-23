package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	benchmarks := []struct {
		name string
		num  int
	}{
		{"10th Fib Number", 10},
		{"20th Fib Number", 20},
		{"30th Fib Number", 30},
		{"50th Fib Number", 50},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibRecurse(bm.num)
			}
		})
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibLoop(bm.num)
			}
		})
	}
}
