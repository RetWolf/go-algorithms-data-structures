package main

import "testing"

func BenchmarkSummationLoop(b *testing.B) {
	benchmarks := []struct {
		name string
		num  int
	}{
		{"Thousand", 1000},
		{"Billion", 1000000000},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sumLoop(bm.num)
			}
		})
	}
}

func BenchmarkSummationFormula(b *testing.B) {
	benchmarks := []struct {
		name string
		num  int
	}{
		{"Thousand", 1000},
		{"Billion", 1000000000},
		{"Quintillion", 1000000000000000000},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sumFormula(bm.num)
			}
		})
	}
}
